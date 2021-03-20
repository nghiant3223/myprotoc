PROTOS_PATH=api/v1/proto
CONTAINER_WORKDIR=/home/docker
DOCKER_IMAGE=myprotoc

gen:
	docker run --rm --volume $(shell pwd):$(CONTAINER_WORKDIR) $(DOCKER_IMAGE) \
		protoc \
		--go_out=. --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=paths=import \
		$(PROTOS_PATH)/*.proto