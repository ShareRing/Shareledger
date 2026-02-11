#!/opt/homebrew/bin/bash
# Initialize a single-node development chain using an exported mainnet genesis
# This script transforms a mainnet genesis for local development by:
# - Replacing all mainnet validators with a new local validator
# - Adding dev accounts with test balances
# - Updating module parameters for development

set -e

# Configuration
CHAIN_ID="${CHAIN_ID:-ShareRing-VoyagerNet}"
MONIKER="${MONIKER:-dev-node}"
HOME_DIR=~/.shareledger
BINARY=~/go/bin/shareledger
DENOM="nshr"
STAKE_DENOM="nshr"

# Path to the exported mainnet genesis
GENESIS_JSON_PATH="${GENESIS_JSON_PATH:-$HOME/genesis.json}"

# Must be >= chain's DefaultPowerReduction to create a validator at genesis.
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

echo "=== Shareledger Development Chain Initialization (From Mainnet Genesis) ==="
echo "Chain ID: $CHAIN_ID"
echo "Moniker: $MONIKER"
echo "Home: $HOME_DIR"
echo "Source Genesis: $GENESIS_JSON_PATH"

# Verify source genesis exists
if [ ! -f "$GENESIS_JSON_PATH" ]; then
    echo "ERROR: Source genesis file not found: $GENESIS_JSON_PATH"
    exit 1
fi

# Wipe the home directory
rm -rf $HOME_DIR
echo "Wiped home directory: $HOME_DIR"

# Initialize the chain (creates config files and directory structure)
echo "Initializing chain skeleton..."
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

# Get validator pubkey for gentx
VALIDATOR_PUBKEY=$($BINARY tendermint show-validator --home "$HOME_DIR")
echo "Validator pubkey: $VALIDATOR_PUBKEY"

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

# Copy mainnet genesis and transform it
GENESIS_FILE="$HOME_DIR/config/genesis.json"
echo "Copying mainnet genesis..."
cp "$GENESIS_JSON_PATH" "$GENESIS_FILE"

echo "Transforming genesis for local development..."

# Step 1: Reset initial_height to 1
echo "  - Resetting initial_height to 1..."
jq '.initial_height = "1"' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 2: Clear validator-related state
echo "  - Clearing staking validators..."
jq '.app_state.staking.validators = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

echo "  - Clearing last_validator_powers..."
jq '.app_state.staking.last_validator_powers = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

echo "  - Resetting last_total_power..."
jq '.app_state.staking.last_total_power = "0"' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

echo "  - Clearing delegations..."
jq '.app_state.staking.delegations = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

echo "  - Clearing unbonding_delegations..."
jq '.app_state.staking.unbonding_delegations = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

echo "  - Clearing redelegations..."
jq '.app_state.staking.redelegations = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Note: The 'pool' field is not used in the current staking module, so we don't add it

# Step 3: Clear distribution validator state
echo "  - Clearing distribution validator state..."
jq '.app_state.distribution.validator_current_rewards = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.distribution.validator_accumulated_commissions = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.distribution.validator_historical_rewards = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.distribution.validator_slash_events = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.distribution.delegator_starting_infos = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.distribution.outstanding_rewards = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.distribution.previous_proposer = ""' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 4: Clear slashing state
echo "  - Clearing slashing state..."
jq '.app_state.slashing.signing_infos = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.slashing.missed_blocks = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 5: Clear genutil gen_txs (will be populated by gentx command)
echo "  - Clearing genutil gen_txs..."
jq '.app_state.genutil.gen_txs = []' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 5.5: Fix denom_metadata with empty name fields (required for bank module validation)
echo "  - Fixing denom_metadata with empty name fields..."
jq '
    .app_state.bank.denom_metadata = [
        .app_state.bank.denom_metadata[] |
        if .name == "" then .name = .display else . end |
        if .symbol == "" then .symbol = .display else . end
    ]
' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 6: Add/update dev account balances
echo "  - Adding dev account balances..."

# Track total new tokens to add to supply (using bc for arbitrary precision)
declare -A NEW_TOKENS
NEW_TOKENS[nshr]="0"
NEW_TOKENS[cent]="0"

# Function to add or update balance entry
add_or_update_balance() {
    local addr="$1"
    local coins="$2"
    local genesis="$3"
    
    # Check if address already exists
    local exists=$(jq --arg addr "$addr" '[.app_state.bank.balances[] | select(.address == $addr)] | length' "$genesis")
    
    # Parse coins
    IFS=',' read -ra COIN_ARRAY <<< "$coins"
    
    if [ "$exists" -gt 0 ]; then
        echo "      (address exists in mainnet, updating balance)"
        # Address exists - update each coin
        for coin in "${COIN_ARRAY[@]}"; do
            local amount=$(echo "$coin" | sed 's/[^0-9]//g')
            local denom=$(echo "$coin" | sed 's/[0-9]//g')
            
            # Get existing amount for this denom (0 if doesn't exist)
            local existing=$(jq -r --arg addr "$addr" --arg denom "$denom" '
                [.app_state.bank.balances[] | select(.address == $addr)][0].coins // [] | 
                map(select(.denom == $denom)) | 
                if length > 0 then .[0].amount else "0" end
            ' "$genesis")
            existing=${existing:-0}
            
            # Calculate new total using bc for precision
            local new_total=$(echo "$existing + $amount" | bc)
            
            # Only track what we're actually adding
            NEW_TOKENS[$denom]=$(echo "${NEW_TOKENS[$denom]:-0} + $amount" | bc)
            
            # Update the balance
            jq --arg addr "$addr" --arg denom "$denom" --arg newamt "$new_total" '
                .app_state.bank.balances = [
                    .app_state.bank.balances[] |
                    if .address == $addr then
                        if (.coins | map(.denom) | index($denom)) then
                            .coins = [.coins[] | if .denom == $denom then .amount = $newamt else . end] | .coins |= sort_by(.denom)
                        else
                            .coins = (.coins + [{"amount": $newamt, "denom": $denom}]) | .coins |= sort_by(.denom)
                        end
                    else
                        .
                    end
                ]
            ' "$genesis" > "$genesis.tmp" && mv "$genesis.tmp" "$genesis"
        done
    else
        # Address doesn't exist - add new entry
        local coins_json="[]"
        for coin in "${COIN_ARRAY[@]}"; do
            local amount=$(echo "$coin" | sed 's/[^0-9]//g')
            local denom=$(echo "$coin" | sed 's/[0-9]//g')
            coins_json=$(echo "$coins_json" | jq --arg amt "$amount" --arg den "$denom" '. + [{"amount": $amt, "denom": $den}]')
            
            # Track new tokens
            NEW_TOKENS[$denom]=$(echo "${NEW_TOKENS[$denom]:-0} + $amount" | bc)
        done
        
        jq --arg addr "$addr" --argjson coins "$coins_json" '
            .app_state.bank.balances += [{"address": $addr, "coins": ($coins | sort_by(.denom))}]
        ' "$genesis" > "$genesis.tmp" && mv "$genesis.tmp" "$genesis"
    fi
}

# Add balances for all config accounts
for name in "${CONFIG_ACCOUNTS[@]}"; do
    echo "    Adding balance for $name (${CONFIG_ADDR[$name]}): ${CONFIG_COINS[$name]}"
    add_or_update_balance "${CONFIG_ADDR[$name]}" "${CONFIG_COINS[$name]}" "$GENESIS_FILE"
done

# Add balances for alice and bob
echo "    Adding balance for alice ($ALICE_ADDRESS): 100000000000${DENOM}"
add_or_update_balance "$ALICE_ADDRESS" "100000000000${DENOM}" "$GENESIS_FILE"
echo "    Adding balance for bob ($BOB_ADDRESS): 100000000000${DENOM}"
add_or_update_balance "$BOB_ADDRESS" "100000000000${DENOM}" "$GENESIS_FILE"

# Step 7: Update bank supply by adding the new tokens
echo "  - Updating bank supply with new tokens..."
echo "    New tokens to add: nshr=${NEW_TOKENS[nshr]}, cent=${NEW_TOKENS[cent]}"

for denom in "${!NEW_TOKENS[@]}"; do
    if [ "${NEW_TOKENS[$denom]}" != "0" ]; then
        # Get current supply for this denom
        current_supply=$(jq -r --arg denom "$denom" '.app_state.bank.supply[] | select(.denom == $denom) | .amount // "0"' "$GENESIS_FILE")
        current_supply=${current_supply:-0}
        
        # Calculate new supply using bc for precision
        new_supply=$(echo "$current_supply + ${NEW_TOKENS[$denom]}" | bc)
        echo "    $denom: $current_supply + ${NEW_TOKENS[$denom]} = $new_supply"
        
        # Update supply
        jq --arg denom "$denom" --arg amt "$new_supply" '
            .app_state.bank.supply = [
                .app_state.bank.supply[] |
                if .denom == $denom then .amount = $amt else . end
            ] |
            if (.app_state.bank.supply | map(.denom) | index($denom) | not) then
                .app_state.bank.supply += [{"denom": $denom, "amount": $amt}]
            else . end |
            .app_state.bank.supply |= sort_by(.denom)
        ' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
    fi
done

# Step 8: Update electoral authority and treasurer
echo "  - Updating electoral authority and treasurer..."
jq --arg authority "${CONFIG_ADDR[authority]}" --arg treasurer "${CONFIG_ADDR[treasurer]}" '
  .app_state.electoral.authority.address = $authority
  | .app_state.electoral.treasurer.address = $treasurer
' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 9: Update distributionx dev_pool_account
CONFIG_DEV_POOL_ACCOUNT="${CONFIG_ADDR[treasurer]}"
echo "  - Updating distributionx dev_pool_account to $CONFIG_DEV_POOL_ACCOUNT..."
jq --arg addr "$CONFIG_DEV_POOL_ACCOUNT" '.app_state.distributionx.params.dev_pool_account = $addr' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 10: Update accStateList role entries
echo "  - Updating electoral accStateList roles..."
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

# Step 11: Shorten gov voting periods for development
echo "  - Shortening gov voting periods..."
jq '.app_state.gov.voting_params.voting_period = "10s"' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
jq '.app_state.gov.deposit_params.max_deposit_period = "10s"' "$GENESIS_FILE" > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"

# Step 12: Add auth accounts for new addresses if they don't exist
echo "  - Ensuring auth accounts exist for dev addresses..."
add_auth_account() {
    local addr="$1"
    local genesis="$2"
    
    # Check if account already exists
    local exists=$(jq --arg addr "$addr" '.app_state.auth.accounts | map(select(.address == $addr or (.base_account.address // "") == $addr)) | length' "$genesis")
    
    if [ "$exists" -eq 0 ]; then
        # Add new BaseAccount
        jq --arg addr "$addr" '
            .app_state.auth.accounts += [{
                "@type": "/cosmos.auth.v1beta1.BaseAccount",
                "address": $addr,
                "pub_key": null,
                "account_number": "0",
                "sequence": "0"
            }]
        ' "$genesis" > "$genesis.tmp" && mv "$genesis.tmp" "$genesis"
        echo "    Added auth account for $addr"
    fi
}

for name in "${CONFIG_ACCOUNTS[@]}"; do
    add_auth_account "${CONFIG_ADDR[$name]}" "$GENESIS_FILE"
done
add_auth_account "$ALICE_ADDRESS" "$GENESIS_FILE"
add_auth_account "$BOB_ADDRESS" "$GENESIS_FILE"

# Clear any old gentxs
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
echo "Source genesis: $GENESIS_JSON_PATH"
echo "Chain ID: $CHAIN_ID"
echo "Initial height: 1 (reset from mainnet)"
echo ""
echo "Accounts (keyring-backend: test):"
for name in "${CONFIG_ACCOUNTS[@]}"; do
    echo "  - $name: ${CONFIG_ADDR[$name]}"
done
echo "  - alice: $ALICE_ADDRESS"
echo "  - bob: $BOB_ADDRESS"
echo ""
echo "Note: Only newly-created alice/bob mnemonics are saved to /tmp/*_key.txt"
echo ""
echo "To start the chain, run:"
echo "  $BINARY start --home $HOME_DIR"

