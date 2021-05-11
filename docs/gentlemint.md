#
IdSigner management is moved to `gentlemint` module

### Enroll id signer
```
./slcli tx gentlemint enroll-id-signer shareledger1s884xxxz4k4es35gk9xtx8gsdvl2axmggfaw2u shareledger1s884xxxz4k4es35gk9xtx8gsdvl2axmggfaw2u --from authority --fees 10shr --yes
```
### Remoke id signer
```
./slcli tx gentlemint revoke-id-signer shareledger1s884xxxz4k4es35gk9xtx8gsdvl2axmggfaw2u shareledger1s884xxxz4k4es35gk9xtx8gsdvl2axmggfaw2u --from authority --fees 10shr --yes
```

## Account type
### Authority
- Enroll SHRP loaders
```
./slcli tx gentlemint enroll-loaders shareledger1c6jwfsh6s7y2xtw0gn7p3ndl4zmhjrkzlqqqnm  --key-seed=./authority-seed.json -y -b=block
```

### Account operator
- Create and revoke id signer, document issuer accounts
### Id signer
- Has right to create id
### Document issuer
- Hash right to issuer document

### SHRP loader
```
./slcli tx gentlemint load-shrp shareledger1c6jwfsh6s7y2xtw0gn7p3ndl4zmhjrkzlqqqnm 1 --key-seed=./authority-seed.json -y -b=block
```

## Transfer coin
### Transfer SHRP
```
./slcli tx gentlemint send-shrp shareledger1cgp90s8r8svjh998uykktrr3wvruun779e29nl 1 --key-seed=./treasurer-seed.json -y -b=block
```
### Transfer SHR
```
./slcli tx gentlemint send-shr shareledger1cgp90s8r8svjh998uykktrr3wvruun779e29nl 1 --key-seed=./treasurer-seed.json -y -b=block
```