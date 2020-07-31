SHELL := /bin/bash

BUILD_DIR := /liminex/services/examples/go-grpc

GIT_TAG := $(shell git rev-parse --short HEAD 2> /dev/null)
BUILD_TAG := $(if $(BUILD_TAG),$(BUILD_TAG),goguardian/examples-go-grpc:$(GIT_TAG))
GOROOT := $(shell go env GOROOT)

print-%: ; @echo $*=$($*)

deps:
	$(info --- Installing dependencies)
	cd /liminex/golang && make install

timezones:
	cp $(GOROOT)/lib/time/zoneinfo.zip $(BUILD_DIR)

binary: timezones deps
	$(info --- Building golang binary)
	cd $(BUILD_DIR) && CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -a -installsuffix cgo -o go-grpc .

docker-build: binary
	$(info +++ Building docker container $(BUILD_TAG))
	cd $(BUILD_DIR) && docker build -t $(BUILD_TAG) .

docker-push: docker-build
	$(info +++ Pushing docker image $(BUILD_TAG))
	cd $(BUILD_DIR) && /liminex/utilities/cish/scripts/docker-push.sh $(BUILD_TAG)

test-unit: deps
	$(info +++ Running unit tests)
	cd $(BUILD_DIR) && go test -v -mod=readonly ./...

test: test-unit

ci: test-unit docker-push
