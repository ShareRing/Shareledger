# Shareledger migrate genesis file
There are differents of the state so that we need to migrate the genesis file.

## v0.0.1 to v1.1.0
Moving data from old version to new version of identity module.

```
./shareledger custom-migrate 0.0.1 ./old-genesis.json ./new-genesis.json ./ouput-genesis.json
```
* `old-genesis.json`: The genesis file that is exported from current blockchain.
* `new-genesis.json`: The genesis file that is exported from new blockchain. This file is just the template.
* `ouput-genesis.json`: output genesis file.

## v1.1.0 to v1.1.1
Decimalization, remove old `identity` data.

```
./shareledger custom-migrate 1.1.0 ./input-genesis.json ./ouput-genesis.json
```
* `input-genesis.json`: The genesis file that is exported from current blockchain.
* `ouput-genesis.json`: output genesis file.

