PACKAGES=$(shell go list ./... | grep -v '/vendor/')
PACKAGES_NOCLITEST=$(shell go list ./... | grep -v '/vendor/' | grep -v bitbucket.org/shareringvn/cosmos-sdk/cmd/gaia/cli_test)
COMMIT_HASH:=$(shell git rev-parse --short HEAD)
BUILD_FLAGS=-ldflags "-X bitbucket.org/shareringvn/cosmos-sdk/version.GitCommit=${COMMIT_HASH}"

PREFIX=./build/env.sh

################################################
# BUILD
################################################


ifeq ($(OS), Windows_NT)
	SUFFIX=".exe"
else
	SUFFIX=""
endif


build:
	${PREFIX} go build $(BUILD_FLAGS) -o build/main$(SUFFIX) cmd/main.go
	${PREFIX} go build $(BUILD_FLAGS) -o build/test$(SUFFIX) cmd/test.go
	${PREFIX} go build $(BUILD_FLAGS) -o build/testAmino$(SUFFIX) cmd/testAmino.go
	${PREFIX} go build $(BUILD_FLAGS) -o build/genesis$(SUFFIX) cmd/genesis/main.go

build_linux:
	${PREFIX} env GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/main_linux$(SUFFIX) cmd/main.go
	${PREFIX} env GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/test_linux$(SUFFIX) cmd/test.go

build_rp:
	${PREFIX} env GOOS=linux GOARCH=arm GOARM=7 go build $(BUILD_FLAGS) -o build/main_linux$(SUFFIX) cmd/main.go
	${PREFIX} env GOOS=linux GOARCH=arm GOARM=7 go build $(BUILD_FLAGS) -o build/test_linux$(SUFFIX) cmd/test.go



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
	@build/env.sh dep ensure -v -update

.PHONY: build build_linux get_vendor_deps get_tools update_tools check_tools




