FROM golang:1.21.5-alpine3.18 AS builder

RUN go version

COPY ./ /makeshort-bot
WORKDIR /makeshort-bot

RUN go mod download
RUN go build -o ./.bin/makeshort-bot ./cmd/telegram-bot/main.go


# Lightweight docker container with binary files
FROM alpine:latest

WORKDIR /app

COPY --from=builder /makeshort-bot/.bin/ ./.bin
COPY --from=builder /makeshort-bot/config/ ./config

CMD ["./.bin/makeshort-bot"]