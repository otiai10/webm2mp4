FROM golang:1.12

LABEL maintainer "otiai10 <otiai10@gmail.com>"

RUN apt-get -qq update
RUN apt-get -y install libav-tools

ENV GO111MODULE=on
ADD . $GOPATH/src/github.com/otiai10/webm2mp4
WORKDIR $GOPATH/src/github.com/otiai10/webm2mp4
RUN go get ./...

CMD $GOPATH/bin/webm2mp4
