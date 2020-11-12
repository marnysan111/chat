FROM golang:latest

RUN mkdir /go/src/chat

ADD . /go/src/chat

WORKDIR /go/src/chat

RUN go get github.com/gin-gonic/gin \
    && go get gorm.io/gorm \
    && go get gorm.io/driver/mysql \
    && go get github.com/gorilla/websocket