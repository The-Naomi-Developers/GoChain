package main

import (
	"bytes"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"gopkg.in/labstack/echo.v4"
)

// Basic screenshot impemtation
func screenshot(c echo.Context) error {
	var delay int = 3
	ctx := c.Get("ctx").(context.Context)
	url := c.QueryParam("url")
	if url == "" {
		return c.String(http.StatusNotAcceptable, `Required parameter "url" is empty.`)
	}
	if c.QueryParam("delay") != "" {
		var err error
		delay, err = strconv.Atoi(c.QueryParam("delay"))
		if err != nil {
			return c.String(http.StatusNotAcceptable, `Datatype for parameter "delay" is invalid.`)
		}
	}

	var buf []byte
	if err := chromedp.Run(ctx, fullScreenshot(url, 90, time.Duration(delay)*time.Second, &buf)); err != nil {
		return err
	}
	reader := bytes.NewReader(buf)

	return c.Stream(http.StatusOK, "image/png", reader)
}
