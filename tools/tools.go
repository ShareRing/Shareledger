//go:build tools

package tools

import (
	_ "github.com/cosmos/gogoproto/protoc-gen-gocosmos"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar"
)
