# Changelog

## [0.2.1] - 2019-03-01

### Added
- Adding Zero Fee for Query
- Cosmos-Wrapper to wrap around cosmos-sdk including BaseApp, Router, Handlers. Update handlers such as FeeAmount, FeeDenom. Handlers for Auth, Bank, Booking and Exchange module
- Bump Shareledger version and suffix
- Map *TDM Address* to GenesisValidator
- *ConvertTDMPubKey* to convert from Shareledger PubKey to TDM one
- Store to map TDMaddress to SHR address


### Updated
- *testnet* subcommand to create *data* dir as required by TDMv0.30.0
- Move CosmosBech32 Prefix to *main* command to automatically initialize across other subcommands
- *Nonce* query using new path for ClientContext
- Update ValidatorSet and EndBlocker to convert ProposerAddress received from Tendermint to SHR validator address
- swap *tendermint/tmlibs* to *tendermint/tendermint/libs*



### Fixed
- *unsafe_reset_all* doesn't remove also *data/priv_validator_state.json*
- *GenPrivateKey* randomly generates seed for PrivateKey
- *show_address* to print Bech32 address

## [0.2.0] - 2019-01-23
### Added
- Upgrade cosmos 0.29.0, tendermint 0.27.4
- Shareledger supports Bech32

### Updated
- *init* command
- Query Command to support Nonce, Check balance

### Fixed
- ABCI repo from Tendermint updated
- Address moved to AccAddress
- Sha3 Updated to golang/crypto library


## [0.1.2] - 2019-01-24
### Added
- Command line commands `send_coins`

### Fixed
- Decimal unmarshal with 0.0
- 

## [0.1.1] - 2019-01-05
### Added
- Command line commands `begin_unbonding`, `complete_unbonding`, `val_dist_info`, `reg_validator`, `show_address`, `show_balance`, `testnet`
- API returns a Tendermint Reponse 
- Update Power and Validator List returned to Tendermint
- Initialize Shareledger with `private_key`
- Update testnet config files generation into separate folders
- Update LooseToken, BondedToken in `genesis` file
- Checking minimum token stake

### Changed
- Update logging statements
- Minimum tokens take

### Removed
- None

## [0.1.0] - 2019-12-05
* Shareledger with dual-tokens SHR and SHRP.



