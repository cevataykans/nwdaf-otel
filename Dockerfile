FROM golang:1.24 AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o nwdaf
RUN ls -lh /app

FROM alpine:latest AS release

WORKDIR /
COPY --from=builder /app/nwdaf /nwdaf
RUN chmod +x /nwdaf

CMD ["/nwdaf"]