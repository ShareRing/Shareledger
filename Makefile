PACKAGES=$(shell go list ./... | grep -v '/vendor/')
PACKAGES_NOCLITEST=$(shell go list ./... | grep -v '/vendor/' | grep -v bitbucket.org/shareringvn/cosmos-sdk/cmd/gaia/cli_test)
COMMIT_HASH:=$(shell git rev-parse --short HEAD)
BUILD_FLAGS=-ldflags "-X bitbucket.org/shareringvn/cosmos-sdk/version.GitCommit=${COMMIT_HASH}"

################################################
# BUILD
################################################


ifeq ($(OS), Windows_NT)
	SUFFIX=".exe"
else
	SUFFIX=""
endif


build:
	./build/env.sh go build $(BUILD_FLAGS) -o build/main$(SUFFIX) cmd/main.go
	./build/env.sh go build $(BUILD_FLAGS) -o build/test$(SUFFIX) cmd/test.go


############################################
### Tools & dependencies
############################################

check_tools:
	cd tools && $(MAKE) check_tools

update_tools:
	cd tools && $(MAKE) update_tools

get_tools:
	cd tools && $(MAKE) get_tools

get_vendor_deps:
	@rm -rf vendor/
	@echo "--> Running dep ensure"
	@build/env.sh dep ensure -v

.PHONY: build




