# syntax=docker/dockerfile:1
# https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/

FROM golang:1.16

# Set destination for COPY
WORKDIR /app

# (参考) https://qiita.com/ebkn/items/0b30bdbf0dae5df73d2e
RUN apt update \
    && apt install -y \
    git \
    gcc \
    musl-dev \
    make \
    wget

# # Download Go modules
COPY app/go/go.mod app/go/go.sum app/go/Makefile ./

RUN go mod download

RUN make install-tools

# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can (optionally) document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 80