.PHONY: proto run

all: run

run:
	go run cmd/main.go

proto:
	protoc \
		--go_out=internal/protobuf/user \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/protobuf/user \
		--go-grpc_opt=paths=source_relative \
		--proto_path=../shared/protobufs \
		../shared/protobufs/user.proto