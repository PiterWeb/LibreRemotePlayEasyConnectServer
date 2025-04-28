# Golang build stage
FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o LibreRemotePlayEasyConnectServer .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/LibreRemotePlayEasyConnectServer .

ENTRYPOINT [ "/app/LibreRemotePlayEasyConnectServer" ]  

EXPOSE 80/tcp