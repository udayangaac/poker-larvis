FROM golang:1.19.2-alpine3.16

RUN mkdir /opt/poker-larvis

ADD . /opt/poker-larvis

WORKDIR /opt/poker-larvis

RUN go build -o bin/poker cmd/poker/poker.go 

ENV PATH="${PATH}:/opt/poker-larvis/bin"

WORKDIR /