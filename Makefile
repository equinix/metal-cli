default: build

build:
	GOOS=linux go build -o bin/linux/packet
	GOOS=darwin go build -o bin/darwin/packet

clean: 
	rm -rf bin/

dd: 
	GENDOCS=true go run main.go
# GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

# export GO15VENDOREXPERIMENT:=1
# export CGO_ENABLED:=0
# export GOARCH:=amd64

# LOCAL_OS:=$(shell uname | tr A-Z a-z)
# GOFILES:=$(shell find . -name '*.go' | grep -v -E '(./vendor)')
# GOPATH_BIN:=$(shell echo ${GOPATH} | awk 'BEGIN { FS = ":" }; { print $1 }')/bin
# LDFLAGS=-X github.com/1and1/oneandone-flex-volume/pkg/version.Version=$(shell $(CURDIR)/build/git-version.sh)

# all: \
# 	fmtcheck \
# 	_output/bin/linux/oneandone-flex-volume \
# 	_output/bin/darwin/oneandone-flex-volume \

# release: \
# 	clean \
# 	check \
# 	_output/release/oneandone-flex-volume.tar.gz \

# install: _output/bin/$(LOCAL_OS)/oneandone-flex-volume
# 	cp $< $(GOPATH_BIN)

# _output/bin/%: $(GOFILES)
# 	mkdir -p $(dir $@)
# 	GOOS=$(word 1, $(subst /, ,$*)) go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $@ github.com/1and1/oneandone-flex-volume/cmd/$(notdir $@)

# _output/release/oneandone-flex-volume.tar.gz: _output/bin/linux/oneandone-flex-volume
# 	mkdir -p $(dir $@)
# 	tar czf $@ -C _output/bin/linux oneandone-flex-volume

# clean:
# 	rm -rf _output

# # vet: vendor-status fmtcheck
# # 	@echo "go vet ."
# # 	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
# # 		echo ""; \
# # 		echo "Vet found suspicious constructs. Please check the reported constructs"; \
# # 		echo "and fix them if necessary before submitting the code for review."; \
# # 		exit 1; \
# # 	fi

# # fmt:
# # 	gofmt -w $(GOFMT_FILES)

# # vendor-status:
# # 	@govendor status

# # fmtcheck:
# # 	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

# # .PHONY: all check clean install release vendor
