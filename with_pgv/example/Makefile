PROTOS_PATH=api/v1/proto
CONTAINER_GOPATH=/home/go
CONTAINER_WORKDIR=/home/docker

gen:
	docker run --rm --volume $(shell pwd):$(CONTAINER_WORKDIR) nghiant3223/tikiprotoc \
		protoc \
		-I . \
		-I $(CONTAINER_GOPATH)/src \
		--go_out=. --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=paths=import \
		-I $(CONTAINER_GOPATH)/src/github.com/envoyproxy/protoc-gen-validate \
		--validate_out="lang=go:." \
		$(PROTOS_PATH)/*.proto
