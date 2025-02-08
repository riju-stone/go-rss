package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel`
}

func UrltoRssFeed(url string) (RSSFeed, error) {
	// http client to fetch the Rss Feed
	rssClient := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := rssClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer response.Body.Close()

	rssData, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeedPayload := RSSFeed{}
	err = xml.Unmarshal(rssData, &rssFeedPayload)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeedPayload, nil
}
