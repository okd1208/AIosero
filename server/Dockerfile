# 開発環境
FROM golang:1.17

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go mod download
RUN go build -o myapp ./api/cmd

CMD ["fresh", "./myapp"]