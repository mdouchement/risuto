FROM golang:1.6-alpine
MAINTAINER mdouchement

ENV go15vendorexperiment 1
ENV APP_VERSION 42
ENV RISUTO_DATABASE /data/tiedot_db

RUN apk upgrade
RUN apk add --update --no-cache git

RUN go get github.com/Masterminds/glide

RUN mkdir -p /go/src/github.com/mdouchement/risuto
WORKDIR /go/src/github.com/mdouchement/risuto

COPY . /go/src/github.com/mdouchement/risuto/
RUN glide install
RUN go build -o /usr/local/bin/risuto risuto.go

EXPOSE 5000
CMD ["risuto"]
