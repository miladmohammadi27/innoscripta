check-buf:
	which buf || (go install github.com/bufbuild/buf/cmd/buf@latest)

install-generators:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

generate-proto: check-buf install-generators
	buf generate --path=internal/api/proto
