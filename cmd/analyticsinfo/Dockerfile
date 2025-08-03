FROM golang:1.24 AS builder

# sqlite CGO dependencies
RUN apt-get update && apt-get install -y libsqlite3-dev gcc

WORKDIR /app

COPY ../.. .
RUN CGO_ENABLED=1 GOOS=linux go build -o cmd/server cmd/analyticsinfo/main.go

FROM alpine:latest AS release

# Last command helps clearing cache, reducing image size
RUN apk add --no-cache sqlite

# Bind volume where DB stores data
VOLUME ["/data"]

WORKDIR /
COPY --from=builder /app/cmd/server /server
RUN chmod +x /server

CMD ["/server"]