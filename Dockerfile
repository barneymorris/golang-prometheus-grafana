FROM golang:1.20.4-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY main.go ./

RUN go run main.go