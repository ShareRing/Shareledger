# ShareLedger #

## Requirements ##

* [`dep`](https://github.com/golang/dep) - package management tool
* [`tendermint`](https://github.com/tendermint/tendermint) - consensus algorithm

## Build ##

* `build/env.sh dep ensure -v` to install neccesary dependant packace to `vendor` folder
* `build/env.sh go build -o build/main cmd/main` to build ShareLedger business logic app

## Execution ##

* Touch `temp.txt`
* Run `Shareledger` by executing: ` ./build/main`
* If the following error happens, remove content of `data` folder located under this repo
* Run `tendermint` by running:  `tendermint --consensus.create_empty_blocks=false --proxy_app tcp://127.0.0.1:46658 --rpc.laddr tcp://0.0.0.0:46656 node`

