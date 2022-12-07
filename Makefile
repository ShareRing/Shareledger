.PHONY: build_linux_amd64 build dbuild dinit dup ddown duprefresh run
VERSION := v0.44
COMMIT := $(shell git log -1 --format='%H')
BUILDDIR := ./build
DOCKER := $(shell which docker)
build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))


ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=shareledger \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=shareledgerd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath
run:
	go run ./cmd/Shareledgerd/main.go start

build_linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o build/shareledger_linux_amd64 ./cmd/Shareledgerd

build:
	go build -mod=readonly $(BUILD_FLAGS) -o build/shareledger ./cmd/Shareledgerd

dbuild:
	docker build -t sharering/shareledger -f ./deploy/docker/Dockerfile . --platform linux/amd64

build-linux:
	echo $(BUILDDIR)
	mkdir -p $(BUILDDIR)
	$(DOCKER) build -f Dockerfile-ubuntu --rm --tag sharering/builder:latest .
	$(DOCKER) create --name shareledger sharering/builder
	$(DOCKER) cp shareledger:/usr/bin/shareledger $(BUILDDIR)/shareledger
	$(DOCKER) rm shareledger

dinit:
	rm -rf ./deploy/testnet && \
	cp -r ./deploy/testnet_config ./deploy/testnet && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node0/config && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node1/config && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node2/config && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node3/config

dup:
	cd ./deploy && \
    docker compose up -d --force-recreate


ddown:
	cd ./deploy && \
    docker compose down

duprefreshall: ddownswap ddown dinit dup dupswap

test:
	go test ./... -v
