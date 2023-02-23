# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/. ./cmd/
COPY ./internal/. ./internal/
RUN go build -o /metal ./cmd/metal

## Image
FROM scratch

WORKDIR /

COPY --from=build /metal /metal

USER nonroot:nonroot

ENTRYPOINT ["/metal"]
