SHELL := /bin/bash

BUILD_DIR := $(HOME)/fullstack-code-challenge/api
GENERATED_DIR := $(HOME)/fullstack-code-challenge/api/proto/gen/go
GW_DIR := $(GENERATED_DIR)/gateways
PROTO_DIR := $(HOME)/fullstack-code-challenge/api/proto
MODULE_DEPS_DIR := ../module-dependencies

print-%: ; @echo $*=$($*)

deps:
	$(info --- Installing dependencies for modules)
	cd $(MODULE_DEPS_DIR) && \
	go install github.com/golang/protobuf/protoc-gen-go && \
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

grpc:
	$(info --- Building gRPC service definitions) \
	mkdir -p "$(GENERATED_DIR)" && \
	protoc \
		-I=$(PROTO_DIR) \
		-I=$(PROTO_DIR)/third-party \
		--go_out=plugins=grpc:"$(GENERATED_DIR)" \
		--swagger_out=logtostderr=true:"$(GENERATED_DIR)" \
		$(PROTO_DIR)/fullstack_code_challenge/v1/fullstack_code_challenge.proto && \
	\
	mkdir -p "$(GW_DIR)" && \
	protoc \
		-I=$(PROTO_DIR) \
		-I=$(PROTO_DIR)/third-party \
		--grpc-gateway_out=logtostderr=true:"$(GENERATED_DIR)" \
		--swagger_out=logtostderr=true:"$(GENERATED_DIR)" \
		$(PROTO_DIR)/fullstack_code_challenge/v1/fullstack_code_challenge.proto

clean:
	rm -rf $(GENERATED_DIR)

install: deps grpc

binary: deps
	$(info --- Building golang binary)
	cd $(BUILD_DIR) && CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -a -installsuffix cgo -o go-grpc .

test-unit: deps
	$(info +++ Running unit tests)
	cd $(BUILD_DIR) && go test -v -mod=readonly ./...

test: test-unit

