# 開発環境
FROM golang:1.17

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
RUN go get github.com/labstack/echo 
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/lib/pq

CMD ["fresh"]