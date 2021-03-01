
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

build_tags= cleveldb

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=shareledger \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=shareledger \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=slcli \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags)"

build:
	go build -tags cleveldb -ldflags '$(ldflags)' -mod=readonly -o build/shareledger ./cmd/shareledger
	go build -tags cleveldb -ldflags '$(ldflags)' -mod=readonly -o build/slcli ./cmd/slcli
	go build -mod=readonly -o build/debug ./cmd/debug

build-docker:
	$(MAKE) -C testnet/docker/ all

clean:
	sudo rm -rf ./build

.PHONY: build clean