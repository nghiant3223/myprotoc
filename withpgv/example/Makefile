PROTOS_PATH=api/v1/proto
CONTAINER_GOPATH=/home/go
CONTAINER_WORKDIR=/home/docker

gen:
	docker run --rm --volume $(shell pwd):$(CONTAINER_WORKDIR) myprotoc-pgv \
		protoc \
		-I . \
		-I $(CONTAINER_GOPATH)/src \
		-I $(CONTAINER_GOPATH)/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=. --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=paths=import \
		--validate_out="lang=go:." \
		$(PROTOS_PATH)/*.proto
