FROM golang:1.8

MAINTAINER otiai10 <otiai10@gmail.com>

RUN apt-get -qq update
RUN apt-get -y install libav-tools

ADD . $GOPATH/src/github.com/otiai10/webm2mp4
WORKDIR $GOPATH/src/github.com/otiai10/webm2mp4
RUN go get ./...

CMD $GOPATH/bin/webm2mp4
