FROM golang:1.14

RUN apt-get update -y && apt-get install -yq --no-install-recommends unzip curl ca-certificates make git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/github.com/grahamjleach/txtboard

WORKDIR /go/src/github.com/grahamjleach/txtboard

ADD ./ /go/src/github.com/grahamjleach/txtboard
