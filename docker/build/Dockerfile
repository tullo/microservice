FROM golang:1.15.6

ENV PB_VER 3.14.0
ENV PB_URL https://github.com/google/protobuf/releases/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_64.zip

RUN apt-get -qq update \
    && apt-get dist-upgrade -qqy \
    && apt-get -qqy install apt-utils apt-transport-https curl git make unzip gettext rsync \
    && apt-get clean all

# install protobuf
RUN mkdir -p /tmp/protoc && \
    curl -L ${PB_URL} > /tmp/protoc/protoc.zip && \
    cd /tmp/protoc && \
    unzip protoc.zip && \
    cp /tmp/protoc/bin/protoc /usr/local/bin && \
    cp -R /tmp/protoc/include/* /usr/local/include && \
    chmod go+rx /usr/local/bin/protoc && \
    cd /tmp && \
    rm -r /tmp/protoc

# Install grpc
RUN GO111MODULE=on go get google.golang.org/grpc

# Install protoc-gen-go
RUN GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go

# Install protoc-gen-twirp
RUN GO111MODULE=on go get github.com/twitchtv/twirp/protoc-gen-twirp

# Install additional twirp outputs
RUN GO111MODULE=on go get github.com/thechriswalker/protoc-gen-twirp_js
RUN GO111MODULE=on go get github.com/elliots/protoc-gen-twirp_swagger

# Install interface scaffolder
RUN GO111MODULE=on go get github.com/josharian/impl

# Install google wire for DI
RUN GO111MODULE=on go get github.com/google/wire/cmd/wire

# Install golint
RUN go get golang.org/x/lint/golint

# GolangCI lint
RUN cd /usr/local && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.0

RUN useradd -m -U -s /bin/bash -u 1000 builder
USER builder
WORKDIR /home/builder
ENV GOPATH /home/builder/go
