FROM golang:1.11

ENV GO111MODULE=on

ADD . /go/src/github.com/kgoralski/go-crud-template
WORKDIR /go/src/github.com/kgoralski/go-crud-template
#RUN go get github.com/kgoralski/go-crud-template

RUN apt-get update && \
    apt-get install -y vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN go install cmd/main.go

ENTRYPOINT scripts/go-app-entrypoint.sh

EXPOSE 8080