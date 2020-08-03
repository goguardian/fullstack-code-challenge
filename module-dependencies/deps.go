package moduledeps

// this file is needed to download and install dependencies that are needed
// to generate the golang proto directory
import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/idubinskiy/fresh"
	_ "github.com/robertkrimen/godocdown/godocdown"
	_ "golang.org/x/lint/golint"
	_ "google.golang.org/genproto"
)
