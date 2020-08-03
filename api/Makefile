SHELL := /bin/bash

BUILD_DIR := $(HOME)/fullstack-code-challenge/api
GENERATED_DIR := $(HOME)/fullstack-code-challenge/generated
GW_DIR := $(GENERATED_DIR)/gateways
PROTO_DIR := $(HOME)/fullstack-code-challenge/proto
MODULE_DEPS_DIR := ../module-dependencies

GIT_TAG := $(shell git rev-parse --short HEAD 2> /dev/null)
BUILD_TAG := $(if $(BUILD_TAG),$(BUILD_TAG),goguardian/fullstack-code-challenge:$(GIT_TAG))
GOROOT := $(shell go env GOROOT)

print-%: ; @echo $*=$($*)

deps:
	$(info --- Installing dependencies for modules)
	cd $(MODULE_DEPS_DIR) && \
	go install github.com/golang/protobuf/protoc-gen-go && \
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

grpc: clean
		$(info --- Building gRPC service definitions) \
		cd $(MODULE_DEPS_DIR) && \
		protoc \
			-I=$(PROTO_DIR) \
			-I=$(PROTO_DIR)/third-party \
			--go_out=plugins=grpc:"$(HOME)/fullstack-code-challenge/generated" \
			--swagger_out=logtostderr=true:"$(GENERATED_DIR)" \
			$(PROTO_DIR)/fullstack-code-challenge.proto && \
		\
		mkdir -p "$(GW_DIR)" && \
		protoc \
			-I=$(PROTO_DIR) \
			-I=$(PROTO_DIR)/third-party \
			--grpc-gateway_out=logtostderr=true:"$(GW_DIR)" \
			--swagger_out=logtostderr=true:"$(GW_DIR)" \
			$(PROTO_DIR)/fullstack-code-challenge.proto

clean:
	rm -rf $(GENERATED_DIR)/*

install: deps grpc

timezones:
	cp $(GOROOT)/lib/time/zoneinfo.zip $(BUILD_DIR)

binary: timezones deps
	$(info --- Building golang binary)
	cd $(BUILD_DIR) && CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -a -installsuffix cgo -o go-grpc .

test-unit: deps
	$(info +++ Running unit tests)
	cd $(BUILD_DIR) && go test -v -mod=readonly ./...

test: test-unit

