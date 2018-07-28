FROM golang:1.10.3

ADD . /go/src/github.com/kgoralski/go-crud-template
WORKDIR /go/src/github.com/kgoralski/go-crud-template
#RUN go get github.com/kgoralski/go-crud-template

RUN apt-get update && \
    apt-get install -y vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    go install github.com/kgoralski/go-crud-template

ENTRYPOINT scripts/go-app-entrypoint.sh

EXPOSE 8080