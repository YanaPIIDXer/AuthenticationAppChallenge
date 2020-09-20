FROM golang:latest

RUN go get github.com/oxequa/realize
RUN go get bitbucket.org/liamstask/goose/cmd/goose && \
    go get github.com/go-sql-driver/mysql

WORKDIR /go/src
ADD . /go/src

CMD realize start
