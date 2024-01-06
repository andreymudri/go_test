# syntax=docker/dockerfile:1

ARG GO_VERSION=1.21.5
FROM golang:${GO_VERSION} AS build

WORKDIR /app

COPY . .

RUN go build -o bin .

# Expose the port that the application listens on.
EXPOSE 8080

# What the container should run when it is started.
ENTRYPOINT [ "/app/bin" ]
