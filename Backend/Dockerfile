FROM golang:latest

RUN go get github.com/oxequa/realize
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/crypto/bcrypt

WORKDIR /go/src
ADD . /go/src

CMD realize start
