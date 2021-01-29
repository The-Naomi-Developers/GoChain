package main

import (
	"bytes"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/labstack/echo"
)

// Basic screenshot impemtation
func screenshot(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return c.String(http.StatusNotAcceptable, `Required parameter "url" is empty.`)
	}

	var buf []byte
	if checkInCache(url) {
		err := readFromCache(url, &buf)
		if err != nil {
			return err
		}
	} else {
		if err := chromedp.Run(ctx, fullScreenshot(url, 90, time.Duration(screenshotDelay)*time.Second, &buf)); err != nil {
			return err
		}
		if err := cache(buf, url); err != nil {
			return err
		}
	}

	reader := bytes.NewReader(buf)

	return c.Stream(http.StatusOK, "image/png", reader)
}
