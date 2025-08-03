FROM golang:1.24 AS builder

WORKDIR /app

COPY ../.. .
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/server cmd/analyticsinfo/main.go

FROM alpine:latest AS release

# Bind volume where DB stores data
VOLUME ["/data"]

WORKDIR /
COPY --from=builder /app/cmd/server /server
RUN chmod +x /server

CMD ["/server"]