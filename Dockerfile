FROM golang:latest

ADD . /go/src/github.com/kgoralski/go-crud-template
WORKDIR /go/src/github.com/kgoralski/go-crud-template
#RUN go get github.com/kgoralski/go-crud-template

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/gorilla/mux
RUN go install github.com/kgoralski/go-crud-template

RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]

ENTRYPOINT /go/bin/go-crud-template

EXPOSE 8080