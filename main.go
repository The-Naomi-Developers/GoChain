package main

import (
	"context"

	"github.com/chromedp/chromedp"
	"gopkg.in/labstack/echo.v4"
)

func main() {
	e := echo.New()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	e.Use(bindContext(ctx))
	e.GET("/", screenshot)

	e.Logger.Fatal(e.Start(":5000"))
}

func bindContext(ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("ctx", ctx)
			return next(c)
		}
	}
}
