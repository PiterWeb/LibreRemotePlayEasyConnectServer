FROM golang:1.24-alpine

WORKDIR /usr/src/app

COPY . .

RUN go build

RUN ./LibreRemotePlayEasyConnectServer