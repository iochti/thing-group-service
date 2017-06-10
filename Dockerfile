FROM golang:1.8
MAINTAINER Luc CHMIELOWSKI <luc.chmielowski@gmail.com>

# Envs
ENV GO15VENDOREXPERIMENT=1

EXPOSE 5003

RUN mkdir -p /go/src/github.com/iochti/thing-group-service
WORKDIR /go/src/github.com/iochti/thing-group-service
COPY . /go/src/github.com/iochti/thing-group-service

RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

RUN rm -rf /go/src/github.com/iochti/thing-group-service

CMD ["thing-group-service"]
