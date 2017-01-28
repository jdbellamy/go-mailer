FROM golang:alpine

RUN mkdir -p /go/src
RUN mkdir -p /go/bin
RUN set -x GOPATH /go
RUN set -x PATH $PATH /go/bin

ADD . /go/src/github.com/jdbellamy/go-rest

WORKDIR /go/src/github.com/jdbellamy/go-rest

RUN apk --no-cache add bash git curl postfix &&\
    go get &&\
    go install

CMD ["go-rest"]