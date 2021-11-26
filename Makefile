# build_linux_arm64:
# 	env GOOS=linux GOARCH=arm64 go build -o build/shareledger_linux_arm64 -mod=readonly cmd/shareledger/main.go

# build_linux_amd64:
# 	env GOOS=linux GOARCH=amd64 go build -o build/shareledger_linux_amd64 -mod=readonly cmd/shareledger/main.go

# build_windows_amd64:
# 	env GOOS=windows GOARCH=amd64 go build -o build/shareledger_windows_amd64.exe -mod=readonly cmd/shareledger/main.go

# build_darwin_amd64:
# 	env GOOS=darwin GOARCH=amd64 go build -o build/shareledger_darwin_amd64 -mod=readonly cmd/shareledger/main.go

# build_all: build_linux_arm64 build_linux_amd64 build_windows_amd64 build_darwin_amd64

# .PHONY: build build_linux_arm64 build_linux_amd64 build_windows_amd64 build_all

# ------------------------------


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

build-macosM1:
	CGO_CFLAGS="-I/opt/homebrew/Cellar/leveldb/1.23/include" CGO_LDFLAGS="-L//opt/homebrew/Cellar/leveldb/1.23/lib" go build -tags leveldb -ldflags '$(ldflags)' -mod=readonly -o build/shareledger ./cmd/shareledger
	CGO_CFLAGS="-I/opt/homebrew/Cellar/leveldb/1.23/include" CGO_LDFLAGS="-L//opt/homebrew/Cellar/leveldb/1.23/lib"  go build -tags leveldb -ldflags '$(ldflags)' -mod=readonly -o build/slcli ./cmd/slcli
	CGO_CFLAGS="-I/opt/homebrew/Cellar/leveldb/1.23/include" CGO_LDFLAGS="-L//opt/homebrew/Cellar/leveldb/1.23/lib"  go build -mod=readonly -o build/debug ./cmd/debug

run_test:
	go test ./tests...
build-docker:
	$(MAKE) -C testnet/docker/ all

clean:
	sudo rm -rf ./build

.PHONY: build clean
