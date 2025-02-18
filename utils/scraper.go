package utils

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/riju-stone/go-rss/internal/database"
	log "github.com/riju-stone/go-rss/logging"
)

func InitRssScraper(db *database.Queries, concurrentReqs int, timeBetReqs time.Duration) {
	log.Info("Started Scraping on %v routines every %v", concurrentReqs, timeBetReqs)

	reqTicket := time.NewTicker(timeBetReqs)
	for ; ; <-reqTicket.C {
		feeds, err := db.FetchLatestFeeds(context.Background(), int32(concurrentReqs))
		if err != nil {
			log.Error("Error fetching latest feeds %v", err)
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go ScrapeFeed(wg, db, feed)
		}
		wg.Wait()
	}
}

func ScrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFetchedFeed(context.Background(), feed.ID)
	if err != nil {
		log.Error("Failed to mark latest fetched feed id=%v", feed.ID)
		return
	}

	rssFeed, err := UrltoRssFeed(feed.Url)
	if err != nil {
		log.Error("Error fetching feed url %v", feed.Url)
	}

	for _, item := range rssFeed.Channel.Item {
		log.Debug("Found Post %v on Feed %v", item.Title, feed.FeedName)

		postDescription := sql.NullString{}
		if item.Description != "" {
			postDescription.String = item.Description
			postDescription.Valid = true
		}

		postPublishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Error("Could not parse published of Feed %v", feed.FeedName)
		}

		_, err = db.CreateNewPost(context.Background(), database.CreateNewPostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: postDescription,
			PublishedAt: postPublishedAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Error("Unable to create post for %v %v", item.Title, err)
		}
	}

	log.Info("Feed %v collected. %v posts found", feed.FeedName, len(rssFeed.Channel.Item))
}
