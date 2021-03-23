RUN apt update && apt upgrade -y
RUN apt install curl unzip git -y

ENV PATH $PATH:/usr/local/go/bin
RUN go version

ENV HOME /home
ENV GOPATH /home/go

ENV PROTOC_VERSION 3.15.6
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

RUN mkdir ${HOME}/docker
ENV DOCKER ${HOME}/docker
WORKDIR ${DOCKER}
