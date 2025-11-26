FROM golang:1.24 AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/exscaler/main.go

FROM alpine:latest AS release

EXPOSE 8080 8081

WORKDIR /
COPY --from=builder /server /server
RUN chmod +x /server

CMD ["/server"]