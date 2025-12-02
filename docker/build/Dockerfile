FROM golang:1.25.5

ENV PB_VER 3.15.8
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
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

# Install protoc-gen-twirp
RUN go install github.com/twitchtv/twirp/protoc-gen-twirp@v7.2.0

# Install additional twirp outputs
RUN go install github.com/thechriswalker/protoc-gen-twirp_js@v0.0.0-20190627152235-0fe8731d4d8f
RUN go install github.com/elliots/protoc-gen-twirp_swagger@v0.0.0-20200502013400-f21ef47d69e3

# Install interface scaffolder
RUN go install github.com/josharian/impl@v1.0.0

# Install google wire for DI
RUN go install github.com/google/wire/cmd/wire@v0.5.0

# Install golint
RUN go install golang.org/x/lint/golint@v0.0.0-20210508222113-6edffad5e616

# GolangCI lint
# RUN cd /usr/local && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.1
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

RUN useradd -m -U -s /bin/bash -u 1000 builder
USER builder
WORKDIR /home/builder
ENV GOPATH /home/builder/go
