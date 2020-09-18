FROM golang:latest

RUN go get github.com/oxequa/realize
RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

CMD realize start
