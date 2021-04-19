FROM golang:1.16-alpine3.13 AS build

ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/wormi4ok/menuplanner

RUN set -xe && apk add --no-cache git

COPY . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8081

HEALTHCHECK --timeout=3s CMD curl -f http://localhost:8081/health || exit 1

ENTRYPOINT CompileDaemon -build="go install" -command="/go/bin/menuplanner"
