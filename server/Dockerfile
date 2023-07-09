# 開発環境
FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ENV GO111MODULE=on

RUN go build -o myapp .

RUN go get github.com/pilu/fresh

CMD ["fresh"]
