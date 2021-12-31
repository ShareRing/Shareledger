.PHONY: build_linux_amd64 build dbuild dinit dup ddown duprefresh

build_linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o build/shareledger_linux_amd64 /cmd/Shareledgerd/main.go

build:
	go build -o build/shareledger ./cmd/Shareledgerd

dbuild:
	docker build -t sharering/shareledger -f ./deploy/docker/Dockerfile.yaml .

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