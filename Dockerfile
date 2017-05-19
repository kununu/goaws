FROM golang:alpine

ADD . /go/src/github.com/kununu/goaws

RUN apk add --no-cache git mercurial && go get -u github.com/kardianos/govendor && apk del git mercurial

RUN cd /go/src/github.com/kununu/goaws && /go/bin/govendor sync

RUN go install github.com/kununu/goaws/app/cmd

ENTRYPOINT ["/go/bin/cmd"]

EXPOSE 4100
