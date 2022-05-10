.PHONY: build_linux_amd64 build dbuild dinit dup ddown duprefresh run

run:
	go run ./cmd/Shareledgerd/main.go start

build_linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o build/shareledger_linux_amd64 ./cmd/Shareledgerd

build:
	go build -o build/shareledger ./cmd/Shareledgerd

dbuild:
	docker build -t sharering/shareledger -f ./deploy/docker/Dockerfile .

dinit:
	rm -rf ./deploy/testnet && \
	cp -r ./deploy/testnet_config ./deploy/testnet && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node0/config && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node1/config && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node2/config && \
	cp ./deploy/testnet/genesis.json ./deploy/testnet/node3/config
dup:
	cd ./deploy && \
    docker-compose up -d --remove-orphans
ddown:
	cd ./deploy && \
    docker-compose down

duprefresh: dinit dup

test:
	go test ./... -v

abigen:
	for i in *.json; do \
      echo "$i" \
    done
#	abigen --pkg sharetoken --abi=/Users/hoai/project/sharering/swap-contract-evm/abi/ShareToken.json --out ./pkg/swap/abi_gen/sharetoken/sharetoken.go &&\
#	abigen --pkg erc20 --abi=/Users/hoai/project/sharering/swap-contract-evm/abi/ERC20.json --out ./pkg/swap/abi_gen/erc20/erc20.go &&\
#	abigen --pkg ierc20 --abi=/Users/hoai/project/sharering/swap-contract-evm/abi/IERC20.json --out ./pkg/swap/abi_gen/ierc20/ierc20.go &&\
#	abigen --pkg ierc20_metadata --abi=/Users/hoai/project/sharering/swap-contract-evm/abi/IERC20Metadata.json --out ./pkg/swap/abi_gen/ierc20_metadata/ierc20_metadata.go &&\
#	abigen --pkg ownable --abi=/Users/hoai/project/sharering/swap-contract-evm/abi/Ownable.json --out ./pkg/swap/abi_gen/ownable/ownable.go &&\
#	abigen --pkg swap --abi=/Users/hoai/project/sharering/swap-contract-evm/abi/Swap.json --out ./pkg/swap/abi_gen/swap/swap.go
MYDIR = /Users/hoai/project/sharering/swap-contract-evm/abi
list: $(MYDIR)/*.json
	for file in $^ ; do \
		echo "Hello" $${file} ; \
	done