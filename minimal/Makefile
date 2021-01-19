PROTOS_PATH=protos/
CONTAINER_WORKDIR=/home/docker

gen:
	docker run --rm --volume $(pwd):$(CONTAINER_WORKDIR) myprotoc
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTOS_PATH)/*.proto