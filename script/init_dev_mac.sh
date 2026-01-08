#!/opt/homebrew/bin/bash  
# Initialize a single-node development chain for Shareledger in macOs 

set -e

# Configuration
CHAIN_ID="${CHAIN_ID:-ShareRing-VoyagerNet}"
MONIKER="${MONIKER:-dev-node}"
HOME_DIR=~/.shareledger
BINARY=~/go/bin/shareledger
DENOM="nshr"
STAKE_DENOM="nshr"

# Must be >= chain's DefaultPowerReduction to create a validator at genesis.
# (Your panic showed DefaultPowerReduction = 1374393862496, so default to a safe higher value.)
STAKE_AMOUNT="${STAKE_AMOUNT:-2000000000000}"

# macOS uses BSD sed, which requires `-i ''` (whereas GNU sed accepts `-i`).
sed_in_place() {
    local expr="$1"
    local file="$2"

    if sed --version >/dev/null 2>&1; then
        # GNU sed
        sed -i -e "$expr" "$file"
    else
        # BSD sed (macOS)
        sed -i '' -e "$expr" "$file"
    fi
}

echo "=== Shareledger Development Chain Initialization ==="
echo "Chain ID: $CHAIN_ID"
echo "Moniker: $MONIKER"
echo "Home: $HOME_DIR"

# wipte the home directory
rm -rf $HOME_DIR
echo "Wiped home directory: $HOME_DIR"

# Initialize the chain
echo "Initializing chain..."
$BINARY init "$MONIKER" --chain-id "$CHAIN_ID" --home "$HOME_DIR" --overwrite

# Accounts (aligned to config.yml)
CONFIG_ACCOUNTS=(validator authority treasurer account-operator idsigner user approver relayer swap_manager doc-issuer)

declare -A CONFIG_MNEMONIC
CONFIG_MNEMONIC[validator]="arrange amateur body cotton slim roof hand blush monkey remove expect rug hazard spoil flag choose tomato end duty nominee wheel cushion into stable"
CONFIG_MNEMONIC[authority]="work like grass pyramid august topic exit wood reunion until retire frown bean cherry stage attack shed oxygen chronic return kiss hint hat future"
CONFIG_MNEMONIC[treasurer]="bird pizza tobacco omit cricket noodle hold wagon opinion shiver scout nature discover almost permit ceiling endless total sight cattle crisp calm popular flat"
CONFIG_MNEMONIC[account-operator]="decline electric decade treat scissors floor fade fade exile swim destroy unusual noodle cabin toy print cover limb old report that balance pelican member"
CONFIG_MNEMONIC[idsigner]="memory mind warfare pull risk math concert address zero speak glimpse outside economy hill boil boss pulp much connect install clip short object tobacco"
CONFIG_MNEMONIC[user]="bright payment cash tomato tragic impulse perfect jacket matter jelly artist pulse will cinnamon erase middle elevator away clinic razor rotate tide unfair trigger"
CONFIG_MNEMONIC[approver]="spend people immense ill property hungry craft corn quote once hand clarify adapt disorder airport balance alley wisdom physical appear debris awake pencil skate"
CONFIG_MNEMONIC[relayer]="indoor donate grid ostrich tree swamp cactus common piano buzz version world second try garage squirrel alert fork december control bind spoon taste essay"
CONFIG_MNEMONIC[swap_manager]="logic fade bike misery female father false speak code immune improve key food enter night timber kick spare amused miss expire bottom walk century"
CONFIG_MNEMONIC[doc-issuer]="loyal siren evoke advice churn behave volcano wood ecology select unusual clock impulse angry scene protect lucky muffin chimney earth type provide taste volcano"

declare -A CONFIG_COINS
CONFIG_COINS[validator]="100000000000000000nshr,100000000000cent"
CONFIG_COINS[authority]="100000000000000000nshr,100000000000cent"
CONFIG_COINS[treasurer]="100000000000000000nshr,100000000000cent"
CONFIG_COINS[account-operator]="100000000000000000nshr,100000000000cent"
CONFIG_COINS[idsigner]="100000000000000000nshr,100000000000cent"
CONFIG_COINS[user]="1000000000000000nshr"
CONFIG_COINS[approver]="100000000000nshr"
CONFIG_COINS[relayer]="100000000000nshr"
CONFIG_COINS[swap_manager]="100000000000nshr"
CONFIG_COINS[doc-issuer]="100000000000nshr"

declare -A CONFIG_ADDR
echo "Recovering config.yml accounts (keyring-backend: test)..."
for name in "${CONFIG_ACCOUNTS[@]}"; do
    if $BINARY keys show "$name" -a --keyring-backend test --home "$HOME_DIR" >/dev/null 2>&1; then
        echo "Key '$name' already exists; skipping recover."
    else
        echo "Recovering key '$name'..."
        printf '%s\n' "${CONFIG_MNEMONIC[$name]}" | $BINARY keys add "$name" --recover --keyring-backend test --home "$HOME_DIR" >/dev/null
    fi
    CONFIG_ADDR[$name]=$($BINARY keys show "$name" -a --keyring-backend test --home "$HOME_DIR")
done

VALIDATOR_ADDRESS="${CONFIG_ADDR[validator]}"
echo "Validator address: $VALIDATOR_ADDRESS"

# Create extra dev accounts (alice/bob)
echo "Creating extra dev accounts (alice/bob)..."
if $BINARY keys show alice -a --keyring-backend test --home "$HOME_DIR" >/dev/null 2>&1; then
    echo "Key 'alice' already exists; skipping creation."
else
    $BINARY keys add alice --keyring-backend test --home "$HOME_DIR" 2>&1 | tee /tmp/alice_key.txt
fi
if $BINARY keys show bob -a --keyring-backend test --home "$HOME_DIR" >/dev/null 2>&1; then
    echo "Key 'bob' already exists; skipping creation."
else
    $BINARY keys add bob --keyring-backend test --home "$HOME_DIR" 2>&1 | tee /tmp/bob_key.txt
fi

ALICE_ADDRESS=$($BINARY keys show alice -a --keyring-backend test --home "$HOME_DIR")
BOB_ADDRESS=$($BINARY keys show bob -a --keyring-backend test --home "$HOME_DIR")

echo "Alice address: $ALICE_ADDRESS"
echo "Bob address: $BOB_ADDRESS"

# Add genesis accounts with initial balances
echo "Adding genesis accounts..."
for name in "${CONFIG_ACCOUNTS[@]}"; do
    $BINARY add-genesis-account "${CONFIG_ADDR[$name]}" "${CONFIG_COINS[$name]}" --home "$HOME_DIR"
done
$BINARY add-genesis-account "$ALICE_ADDRESS" "100000000000${DENOM}" --home "$HOME_DIR"
$BINARY add-genesis-account "$BOB_ADDRESS" "100000000000${DENOM}" --home "$HOME_DIR"

# Update genesis.json for required module params
GENESIS_FILE="$HOME_DIR/config/genesis.json"

# Align distributionx + electoral module params to recovered config.yml addresses
CONFIG_DEV_POOL_ACCOUNT="shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr"
DEV_POOL_ACCOUNT_ADDR=""
DEV_POOL_ACCOUNT_NAME=""
for name in "${CONFIG_ACCOUNTS[@]}"; do
    if [ "${CONFIG_ADDR[$name]}" = "$CONFIG_DEV_POOL_ACCOUNT" ]; then
        DEV_POOL_ACCOUNT_ADDR="${CONFIG_ADDR[$name]}"
        DEV_POOL_ACCOUNT_NAME="$name"
        break
    fi
done
if [ -z "$DEV_POOL_ACCOUNT_ADDR" ]; then
    DEV_POOL_ACCOUNT_ADDR="$VALIDATOR_ADDRESS"
    echo "WARN: config.yml dev_pool_account ($CONFIG_DEV_POOL_ACCOUNT) did not match any recovered key; using validator ($DEV_POOL_ACCOUNT_ADDR)"
else
    echo "Using dev_pool_account from recovered key '$DEV_POOL_ACCOUNT_NAME' ($DEV_POOL_ACCOUNT_ADDR)"
fi

echo "Updating genesis.json with dev_pool_account..."
jq --arg addr "$DEV_POOL_ACCOUNT_ADDR" '.app_state.distributionx.params.dev_pool_account = $addr' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

echo "Aligning electoral authority/treasurer addresses..."
jq --arg authority "${CONFIG_ADDR[authority]}" --arg treasurer "${CONFIG_ADDR[treasurer]}" '
  .app_state.electoral.authority.address = $authority
  | .app_state.electoral.treasurer.address = $treasurer
' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Ensure electoral accStateList role entries exist (required for various module permissions)
echo "Ensuring electoral accStateList role entries..."
jq \
  --arg voter "$VALIDATOR_ADDRESS" \
  --arg approver "${CONFIG_ADDR[approver]}" \
  --arg relayer "${CONFIG_ADDR[relayer]}" \
  --arg swapManager "${CONFIG_ADDR[swap_manager]}" \
  --arg accop "${CONFIG_ADDR[account-operator]}" \
  --arg docIssuer "${CONFIG_ADDR[doc-issuer]}" \
  --arg idsigner "${CONFIG_ADDR[idsigner]}" \
  --arg shrploader "$VALIDATOR_ADDRESS" \
  '
  (.app_state.electoral.accStateList //= [])
  | .app_state.electoral.accStateList += [
      {"key": ("voter" + $voter), "address": $voter, "status": "active"},
      {"key": ("approver" + $approver), "address": $approver, "status": "active"},
      {"key": ("relayer" + $relayer), "address": $relayer, "status": "active"},
      {"key": ("swapManager" + $swapManager), "address": $swapManager, "status": "active"},
      {"key": ("accop" + $accop), "address": $accop, "status": "active"},
      {"key": ("docIssuer" + $docIssuer), "address": $docIssuer, "status": "active"},
      {"key": ("idsigner" + $idsigner), "address": $idsigner, "status": "active"},
      {"key": ("shrploader" + $shrploader), "address": $shrploader, "status": "active"}
    ]
  | .app_state.electoral.accStateList |= (unique_by(.key))
  ' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

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

# Update gov voting_params voting_period to 100s
echo "Updating gov voting_params voting_period to 100s..."
jq '.app_state.gov.voting_params.voting_period = "10s"' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.gov.deposit_params.max_deposit_period = "10s"' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Clear any old gentxs so this script is re-runnable (cosmos-sdk gentx will not overwrite files)
GENTX_DIR="$HOME_DIR/config/gentx"
if [ -d "$GENTX_DIR" ]; then
    EXISTING_GENTXS=$(ls -1 "$GENTX_DIR"/gentx-*.json 2>/dev/null || true)
    if [ -n "$EXISTING_GENTXS" ]; then
        echo "Found existing gentx files; removing to allow re-run:"
        echo "$EXISTING_GENTXS"
        rm -f "$GENTX_DIR"/gentx-*.json
    fi
else
    mkdir -p "$GENTX_DIR"
fi

# Create gentx for validator
echo "Creating genesis transaction..."
echo "Gentx self-delegation: ${STAKE_AMOUNT}${STAKE_DENOM}"
$BINARY gentx validator "${STAKE_AMOUNT}${STAKE_DENOM}" \
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
    sed_in_place 's/enable = false/enable = true/g' "$APP_TOML"
    # Enable swagger
    sed_in_place 's/swagger = false/swagger = true/g' "$APP_TOML"
    # Set minimum gas prices for dev
    sed_in_place 's/^minimum-gas-prices *=.*/minimum-gas-prices = "0.025nshr"/' "$APP_TOML"
    echo "Updated minimum-gas-prices in app.toml to:"
    grep '^minimum-gas-prices' "$APP_TOML" || echo "minimum-gas-prices not found in app.toml"
    # Enable unsafe CORS for development
    sed_in_place 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' "$APP_TOML"
fi

# Configure config.toml for development
CONFIG_TOML="$HOME_DIR/config/config.toml"
if [ -f "$CONFIG_TOML" ]; then
    echo "Configuring config.toml..."
    # Allow all origins for RPC
    sed_in_place 's/cors_allowed_origins = \[\]/cors_allowed_origins = ["*"]/g' "$CONFIG_TOML"
    # Bind RPC to all interfaces
    sed_in_place 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG_TOML"
fi

echo ""
echo "=== Initialization Complete ==="
echo ""
echo "Accounts (keyring-backend: test):"
for name in "${CONFIG_ACCOUNTS[@]}"; do
    echo "  - $name: ${CONFIG_ADDR[$name]}"
done
echo "  - alice: $ALICE_ADDRESS"
echo "  - bob: $BOB_ADDRESS"
echo ""
echo "Note: Only newly-created alice/bob mnemonics are saved to /tmp/*_key.txt"

