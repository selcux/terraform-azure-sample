.PHONY: protos

protos:
	protoc -I protobuf/ protobuf/*.proto --go_out=. --go-grpc_out=.

run:
	go run service/web/cmd
