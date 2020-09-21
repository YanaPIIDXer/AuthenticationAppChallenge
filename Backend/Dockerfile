FROM golang:latest

RUN go get -u \
    github.com/oxequa/realize \
    bitbucket.org/liamstask/goose/cmd/goose \
    github.com/go-sql-driver/mysql \
    golang.org/x/crypto/bcrypt

WORKDIR /go/src/authapp
ADD . /go/src/authapp

CMD realize start
