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
### Account operator
- Create and revoke id signer, document issuer accounts
### Id signer
- Has right to create id
### Document issuer
- Hash right to issuer document