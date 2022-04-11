# syntax=docker/dockerfile:1

FROM golang:1.18-alpine3.15 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o ./server

FROM alpine:3.15 AS release

WORKDIR /app

ENV PORT=5000

COPY --from=build /app/server ./
COPY .env ./
COPY public/ ./public

EXPOSE ${PORT}

ENTRYPOINT [ "/app/server" ]