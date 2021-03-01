# Shareledger migrate genesis file
There are differents of the state so that we need to migrate the genesis file.

### Migrate
```
shareledger custom-migrate <from version> <old genesis file> <new genesis file> <result genesis file>
```
```
./shareledger custom-migrate 0.0.1 ../tests/genesis-v0.0.1.json ../tests/genesis.v0.0.2.json gen.json

```
