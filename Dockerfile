FROM golang:latest as builder
WORKDIR /go/src/naomi/gochain
COPY ./ ./
RUN go get -u github.com/labstack/echo github.com/chromedp/chromedp github.com/dgrijalva/jwt-go golang.org/x/time/rate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gochain

FROM alpine:latest as runtime
WORKDIR /app
RUN apk update && apk add chromium
COPY --from=builder /go/src/naomi/gochain/gochain ./
RUN chmod +x /app/gochain
CMD ["/app/gochain"]
