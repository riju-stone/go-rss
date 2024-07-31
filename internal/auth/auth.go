package auth

import (
	"errors"
	"net/http"
	"strings"

	log "github.com/riju-stone/go-rss/logging"
)

/*
* Function extracts the API Key from the request header
* Header Format Authorization: apikey-<actual api key>
 */
func AuthenticateApiKey(headers http.Header) (string, error) {
	authVal := headers.Get("Authorization")
	if authVal == "" {
		log.Error("user not authenticated - auth info not found")
		return "", errors.New("user not authenticated")
	}

	authParsedValue := strings.Split(authVal, "-")
	if len(authParsedValue) != 2 {
		log.Error("corrupted auth header - failed to authenticate")
		return "", errors.New("failed to authenticate")
	}

	if authParsedValue[0] != "apikey" {
		log.Error("corrupted auth header - failed to authenticate")
		return "", errors.New("failed to authenticate")
	}

	return authParsedValue[1], nil
}
