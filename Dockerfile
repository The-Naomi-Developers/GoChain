FROM golang:latest as builder
WORKDIR /go/src/naomi/gochain
COPY ./ ./
RUN go get -u gopkg.in/labstack/echo.v4
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gochain

FROM alpine:latest as runtime
WORKDIR /app
COPY --from=builder /go/src/naomi/gochain/gochain ./
RUN chmod +x /app/gochain
CMD ["/app/gochain"]
