# Shareledger #

Shareledger is a custom-designed distributed blockchain with [Tendermint](https://github.com/tendermint/tendermint) as a Byzantine-Fault Tolerant State Machine. ShareLedger provides essential building blocks for any rental/sharing services on top of it.

ShareLedger currently ultilizes a dual-token mechanism, SharePay (SHRP) and ShareToken (SHR). The former provides a stable currency for any additional services running on top of ShareLedger while the latter acts as an ultility token.


## Install Shareledger ##

The fatest and easiest way to install `Shareledger` is to run our os-specfic application which guides you through four steps to setup and run a MasterNode. [TO BE RELEASED](https://sharering.network)


### From Binary
To download pre-built binaries, see the [releases page](https://github.com/sharering/shareledger/releases).

### From Source

#### Requirements ##

* [`go`](https://golang.org/doc/install) - compile tool. Version >=1.15
* [`make`](https://www.gnu.org/software/make/) -  compile tool


#### Get Source Code

```
git clone https://github.com/sharering/shareledger.git
cd shareledger
```

#### Compile (_Work on Linux only_)
```
make build
```

This will compile ShareLedger and put the binary in `./build`


# Run

* To start a ShareLedger node
```
./shareledger init
./shareledger start
```
* Start single node for testing purpose: [start single node](./docs/start-single-node.md)


    Run `./build/shareledger -h` _to get more detailed information on how to execute ShareLedger_


---
## Notes
* Get code from branch feature/upgrade-2019-10-22