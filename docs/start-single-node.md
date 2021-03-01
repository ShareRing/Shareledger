# Start single node for test

```
#!/usr/bin/env bash

rm -rf ~/.shareledger
rm -rf ~/.slcli

./shareledger init test --chain-id=testing

./slcli config output json
./slcli config indent true
./slcli config trust-node true
./slcli config chain-id testing
./slcli config keyring-backend test

echo "valve wash okay biology tissue term fire cross solid sword bulb right team enact raven rare frog repeat dust when zebra focus task voice" |./slcli keys add authority --recover --index 0
echo "valve wash okay biology tissue term fire cross solid sword bulb right team enact raven rare frog repeat dust when zebra focus task voice" |./slcli keys add treasurer --recover --index 1
echo "valve wash okay biology tissue term fire cross solid sword bulb right team enact raven rare frog repeat dust when zebra focus task voice" |./slcli keys add validator --recover --index 2
echo "valve wash okay biology tissue term fire cross solid sword bulb right team enact raven rare frog repeat dust when zebra focus task voice" |./slcli keys add account-operator --recover --index 3
echo "valve wash okay biology tissue term fire cross solid sword bulb right team enact raven rare frog repeat dust when zebra focus task voice" |./slcli keys add issuer --recover --index 4
echo "valve wash okay biology tissue term fire cross solid sword bulb right team enact raven rare frog repeat dust when zebra focus task voice" |./slcli keys add idsigner --recover --index 5

./shareledger add-genesis-authority $(./slcli keys show authority -a)
./shareledger add-genesis-treasurer $(./slcli keys show treasurer -a)
./shareledger add-genesis-validator $(./slcli keys show validator -a)
./shareledger add-genesis-account-operator $(./slcli keys show account-operator -a)

./shareledger add-genesis-account $(./slcli keys show authority -a) 1000000shr,1000000shrp
./shareledger add-genesis-account $(./slcli keys show treasurer -a) 1000000shr,1000000shrp
./shareledger add-genesis-account $(./slcli keys show validator -a) 1000000shr,1000000shrp
./shareledger add-genesis-account $(./slcli keys show account-operator -a) 1000000shr,1000000shrp
./shareledger add-genesis-account $(./slcli keys show issuer -a) 1000000shr,1000000shrp
./shareledger add-genesis-account $(./slcli keys show idsigner -a) 1000000shr,1000000shrp


./shareledger gentx --name "validator" --keyring-backend test --amount 1000000shr

echo "Collecting genesis txs..."
./shareledger collect-gentxs

echo "Validating genesis file..."
./shareledger validate-genesis

./shareledger start
```