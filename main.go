package main

import (
	"context"

	"github.com/chromedp/chromedp"
	"gopkg.in/labstack/echo.v4"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func main() {
	e := echo.New()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	e.Use(bindContext(ctx))
	e.GET("/", screenshot)

	e.Logger.Fatal(e.Start(":5000"))
}
