FROM golang:alpine

ADD . /go/src/github.com/kununu/goaws

RUN go install github.com/kununu/goaws

ENTRYPOINT ["/go/bin/goaws"]

EXPOSE 4100
