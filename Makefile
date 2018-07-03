EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=packet
VERSION=0.0.1
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64 386

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: generate-docs
	go fmt ./...
build:
	go build ${LDFLAGS} -o bin/${BINARY}

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/$(BINARY)-$(GOOS)-$(GOARCH))))

clean: 
	rm -rf bin/
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

clean-docs:
	rm -rf docs/	

install:
	go install ${LDFLAGS}

generate-docs: clean-docs
	mkdir -p docs
	GENDOCS=true go run main.go