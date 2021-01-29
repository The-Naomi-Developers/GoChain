package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

// Cache screenshot
func cache(buf []byte, url string) error {
	urlEnc := base64.URLEncoding.EncodeToString([]byte(url))
	if err := ioutil.WriteFile(cacheDir+urlEnc+".png", buf, 0o644); err != nil {
		return err
	}
	fmt.Printf("[INFO] Successfully cached screenshot for URL '%s'", url)
	return nil
}

// Check screenshot in cache
func checkInCache(url string) bool {
	urlEnc := base64.URLEncoding.EncodeToString([]byte(url))
	if _, err := os.Stat(cacheDir + urlEnc + ".png"); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// Get cached screenshot
func readFromCache(url string, res *[]byte) error {
	var err error

	urlEnc := base64.URLEncoding.EncodeToString([]byte(url))

	*res, err = ioutil.ReadFile(cacheDir + urlEnc + ".png")

	return err
}
