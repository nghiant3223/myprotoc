PROTOS_PATH=api/v1/proto
CONTAINER_WORKDIR=/home/docker

gen:
	docker run --rm --volume $(pwd):$(CONTAINER_WORKDIR) myprotoc
	protoc \
		--go_out=. --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=paths=import \
		$(PROTOS_PATH)/*.proto