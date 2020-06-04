package core

import (
	"fmt"
	"log"
	"math/rand"
	urlPkg "net/url"
	"strings"
)

// AddPage adds a recond of shorten url.
// It takes the original url and returns the short URL and error.
func AddPage(rawurl string) (shorturl string, err error) {
	// Validate.
	// 1) Only support (http, https) scheme.
	// 2) Do not support relative path.
	validURL, err := parseRawURL(rawurl)
	if err != nil {
		return "", err
	}

	// Generate the hash for shorten ID with prefix '_'
	sid := "_" + getRandomAlphabetString(10)

	// TODO: Check collison and write the mapping to DB.
	return sid, nil
}

func parseRawURL(rawurl string) (string, error) {
	u, err := urlPkg.ParseRequestURI(rawurl)
	if err != nil {
		log.Printf("Fail to parse %s: err = %s", rawurl, err)
		return "", err
	}
	if !getSupportedScheme()[u.Scheme] {
		err = fmt.Errorf("Unsupported scheme: %s", u.Scheme)
		log.Printf(err.Error())
		return "", err
	}
	return rawurl, nil
}

const alphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func getRandomAlphabetString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		seed := rand.Int()
		sb.WriteByte(alphabets[seed%52])
	}
	return sb.String()
}

func getSupportedScheme() map[string]bool {
	return map[string]bool{
		"http":  true,
		"https": true,
	}
}
