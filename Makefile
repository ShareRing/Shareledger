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

dupswap:
	cd ./deploy && \
	docker-compose -f docker-compose-relayer.yaml up -d

ddownswap:
	cd ./deploy && \
	docker-compose -f docker-compose-relayer.yaml down

ddown:
	cd ./deploy && \
    docker-compose down

duprefreshall: ddownswap ddown dinit dup dupswap

test:
	go test ./... -v

MYDIR = /Users/hoai/project/sharering/swap-contract-evm/abi
list: $(MYDIR)/*.json
	for file in $^ ; do \
		echo "Hello" $${file} ; \
	done