##
## Build
##

FROM golang:1.23.5-alpine3.21 AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /docker-cronjob-be ./cmd/*.go

##
## Deploy
##

FROM alpine:3.21.2

WORKDIR /

COPY --from=build /docker-cronjob-be /docker-cronjob-be
COPY .env .env

EXPOSE 8080

ENTRYPOINT ["/docker-cronjob-be"]