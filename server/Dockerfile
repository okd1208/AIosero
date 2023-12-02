# 開発環境
FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apt-get update && apt-get install -y postgresql-client

COPY . .
ENV GO111MODULE=on

RUN go build -o myapp .

RUN go install github.com/cosmtrek/air@latest
