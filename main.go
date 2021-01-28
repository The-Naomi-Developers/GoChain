package main

import (
	"context"
	"os"
	"strconv"

	"github.com/chromedp/chromedp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/time/rate"
)

var (
	// Internal variables
	ctx    context.Context
	cx     context.Context
	cancel context.CancelFunc
	// User-defined variables
	cacheDir        string = "./cache"
	port            string = ":5000"
	cooldown        int    = 2 // seconds
	screenshotDelay int    = 2 // seconds
	proxyServer     string
)

func main() {
	e := echo.New()

	Init()

	if proxyServer != "" {

		o := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ProxyServer(proxyServer),
		)
		cx, cancel = chromedp.NewExecAllocator(context.Background(), o...)
		defer cancel()
	}

	ctx, cancel = chromedp.NewContext(cx)
	defer cancel()

	rateLimit(e)

	e.GET("/", screenshot)

	e.Logger.Fatal(
		e.Start(port),
	)
}

func rateLimit(e *echo.Echo) {
	// Allow only 30 requests per second per IP (i.e. 1 request per 2 seconds )
	var ratelimit int = 60 / cooldown
	var limiterStore = middleware.NewRateLimiterMemoryStore(rate.Limit(ratelimit))

	limiter := middleware.RateLimiter(limiterStore)

	e.Use(limiter)
}

// Init configuration
func Init() {
	// TODO add non-cache mode
	os.RemoveAll(cacheDir)
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, 0755)
	}
	if _, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		port = ":" + os.Getenv("PORT")
	}
	if cd, err := strconv.Atoi(os.Getenv("COOLDOWN")); err == nil {
		cooldown = cd
	}
	if dl, err := strconv.Atoi(os.Getenv("WAIT_BEFORE_SCREENSHOT")); err == nil {
		screenshotDelay = dl
	}
	proxyServer = os.Getenv("PROXY")
}
