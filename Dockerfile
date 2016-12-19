FROM golang:1.7.4-alpine
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>

RUN set -x && apk add --no-cache bash git make zip
RUN go get -u github.com/kardianos/govendor

ADD . /go/src/github.com/yamamoto-febc/terraform-provider-arukas

WORKDIR /go/src/github.com/yamamoto-febc/terraform-provider-arukas
CMD ["make"]
