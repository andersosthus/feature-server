.PHONY: all

protos:
	protoc --go_out=./pkg --go-grpc_out=./pkg --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative feature-server-proto/feature.proto