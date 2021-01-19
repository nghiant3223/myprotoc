FROM ubuntu:20.04

RUN apt update && apt upgrade -y
RUN apt install wget tar curl unzip git -y

ENV GO_VERSION 1.15.6
RUN wget -q https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin 
RUN go version

ENV HOME /home
ENV GOPATH /home/go

ENV PROTOC_VERSION 3.13.0
RUN PB_REL="https://github.com/protocolbuffers/protobuf/releases" \
    && curl -s -LO ${PB_REL}/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
    && unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d ${HOME}/.local
ENV PATH $PATH:${HOME}/.local/bin
RUN protoc --version

RUN GO111MODULE=on go get \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc 
ENV PATH $PATH:${GOPATH}/bin
RUN which protoc-gen-go
RUN which protoc-gen-go-grpc

RUN mkdir ${HOME}/docker
ENV DOCKER ${HOME}/docker
WORKDIR ${DOCKER}
