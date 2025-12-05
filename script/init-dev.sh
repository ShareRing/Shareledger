#!/bin/bash
# Initialize a single-node development chain for Shareledger
# This script is run automatically on first startup

set -e

# Configuration
CHAIN_ID="${CHAIN_ID:-shareledger-dev}"
MONIKER="${MONIKER:-dev-node}"
HOME_DIR="/root/.shareledger-dev"
BINARY="./tmp/shareledger"
DENOM="nshr"
STAKE_DENOM="nshr"

echo "=== Shareledger Development Chain Initialization ==="
echo "Chain ID: $CHAIN_ID"
echo "Moniker: $MONIKER"
echo "Home: $HOME_DIR"

# Build the binary first if it doesn't exist
if [ ! -f "$BINARY" ]; then
    echo "Building shareledger binary..."
    go build -o "$BINARY" ./cmd/Shareledgerd
fi

# Initialize the chain
echo "Initializing chain..."
$BINARY init "$MONIKER" --chain-id "$CHAIN_ID" --home "$HOME_DIR"

# Create validator key
echo "Creating validator key..."
$BINARY keys add validator --keyring-backend test --home "$HOME_DIR" 2>&1 | tee /tmp/validator_key.txt

# Get validator address
VALIDATOR_ADDRESS=$($BINARY keys show validator -a --keyring-backend test --home "$HOME_DIR")
echo "Validator address: $VALIDATOR_ADDRESS"

# Create test accounts
echo "Creating test accounts..."
$BINARY keys add alice --keyring-backend test --home "$HOME_DIR" 2>&1 | tee /tmp/alice_key.txt
$BINARY keys add bob --keyring-backend test --home "$HOME_DIR" 2>&1 | tee /tmp/bob_key.txt

ALICE_ADDRESS=$($BINARY keys show alice -a --keyring-backend test --home "$HOME_DIR")
BOB_ADDRESS=$($BINARY keys show bob -a --keyring-backend test --home "$HOME_DIR")

echo "Alice address: $ALICE_ADDRESS"
echo "Bob address: $BOB_ADDRESS"

# Add genesis accounts with initial balances
echo "Adding genesis accounts..."
$BINARY add-genesis-account "$VALIDATOR_ADDRESS" "1000000000000${DENOM}" --home "$HOME_DIR"
$BINARY add-genesis-account "$ALICE_ADDRESS" "100000000000${DENOM}" --home "$HOME_DIR"
$BINARY add-genesis-account "$BOB_ADDRESS" "100000000000${DENOM}" --home "$HOME_DIR"

# Update genesis.json for required module params
GENESIS_FILE="$HOME_DIR/config/genesis.json"

# Set dev_pool_account for distributionx module
echo "Updating genesis.json with dev_pool_account..."
jq --arg addr "$VALIDATOR_ADDRESS" '.app_state.distributionx.params.dev_pool_account = $addr' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Add validator as voter in electoral module (required for MsgCreateValidator)
echo "Adding validator as voter in electoral module..."
jq --arg addr "$VALIDATOR_ADDRESS" '.app_state.electoral.accStateList += [{"key": ("voter" + $addr), "address": $addr, "status": "active"}]' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Update staking bond_denom from "stake" to "nshr"
echo "Updating staking bond_denom to nshr..."
jq --arg denom "$DENOM" '.app_state.staking.params.bond_denom = $denom' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Update crisis constant_fee denom to nshr
echo "Updating crisis constant_fee denom to nshr..."
jq --arg denom "$DENOM" '.app_state.crisis.constant_fee.denom = $denom' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Update mint denom to nshr
echo "Updating mint denom to nshr..."
jq --arg denom "$DENOM" '.app_state.mint.params.mint_denom = $denom' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Update gov min_deposit denom to nshr
echo "Updating gov min_deposit denom to nshr..."
jq --arg denom "$DENOM" '.app_state.gov.deposit_params.min_deposit[0].denom = $denom' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Create gentx for validator
echo "Creating genesis transaction..."
$BINARY gentx validator "500000000000${STAKE_DENOM}" \
    --chain-id "$CHAIN_ID" \
    --keyring-backend test \
    --home "$HOME_DIR" \
    --moniker "$MONIKER"

# Collect gentxs
echo "Collecting genesis transactions..."
$BINARY collect-gentxs --home "$HOME_DIR"

# Validate genesis
echo "Validating genesis..."
$BINARY validate-genesis --home "$HOME_DIR"

# Configure app.toml for development
APP_TOML="$HOME_DIR/config/app.toml"
if [ -f "$APP_TOML" ]; then
    echo "Configuring app.toml..."
    # Enable API
    sed -i 's/enable = false/enable = true/g' "$APP_TOML"
    # Enable swagger
    sed -i 's/swagger = false/swagger = true/g' "$APP_TOML"
    # Set minimum gas prices for dev
    sed -i 's/^minimum-gas-prices *=.*/minimum-gas-prices = "0.025nshr"/' "$APP_TOML"
    echo "Updated minimum-gas-prices in app.toml to:"
    grep '^minimum-gas-prices' "$APP_TOML" || echo "minimum-gas-prices not found in app.toml"
    # Enable unsafe CORS for development
    sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' "$APP_TOML"
fi

# Configure config.toml for development
CONFIG_TOML="$HOME_DIR/config/config.toml"
if [ -f "$CONFIG_TOML" ]; then
    echo "Configuring config.toml..."
    # Allow all origins for RPC
    sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = ["*"]/g' "$CONFIG_TOML"
    # Bind RPC to all interfaces
    sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG_TOML"
fi

echo ""
echo "=== Initialization Complete ==="
echo ""
echo "Test accounts created (keyring-backend: test):"
echo "  - validator: $VALIDATOR_ADDRESS"
echo "  - alice: $ALICE_ADDRESS"
echo "  - bob: $BOB_ADDRESS"
echo ""
echo "Key mnemonics saved to /tmp/*_key.txt"
echo ""
echo "Chain will start automatically with Air hot reload."
echo "Edit any .go file to trigger a rebuild!"

