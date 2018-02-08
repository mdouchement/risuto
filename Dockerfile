# build stage
FROM golang:1.9-alpine as build-env
MAINTAINER mdouchement

RUN apk upgrade
RUN apk add --update --no-cache git

RUN go get github.com/Masterminds/glide
RUN go get github.com/gobuffalo/packr/packr
RUN go get github.com/goreleaser/goreleaser

RUN mkdir -p /go/src/github.com/mdouchement/risuto
WORKDIR /go/src/github.com/mdouchement/risuto

COPY . /go/src/github.com/mdouchement/risuto/
# Dependencies
RUN glide install
# Download static assets
RUN go run risuto.go fetch --min
# Build assets
RUN packr -z
# Packr fix until the filename can be specified/prefix (packr init func must be executed first).
RUN mv web/web-packr.go web/assets-packr.go
# Go build
RUN ./build.sh


# final stage
FROM alpine:3.5
MAINTAINER mdouchement

ENV ECHO_ENV production
ENV RISUTO_DATABASE /data/tiedot_db

COPY --from=build-env /go/src/github.com/mdouchement/risuto/dist/linuxamd64/risuto /usr/local/bin/

EXPOSE 5000
CMD ["risuto", "server", "-p", "5000"]
