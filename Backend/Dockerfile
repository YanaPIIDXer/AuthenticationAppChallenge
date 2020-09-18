FROM golang:latest

RUN go get github.com/oxequa/realize

RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

CMD realize start
