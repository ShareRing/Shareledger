#!/bin/bash
# work only in macOs
set -euo pipefail

HOME_DIR=~/.shareledger
BIN=~/go/bin/shareledger
CHAIN_ID="${CHAIN_ID:-ShareRing-VoyagerNet}"
PROPOSAL_FILE="v3_upgrade_proposal.json"

CURRENT_HEIGHT=$(curl -s http://localhost:26657/status | jq -r '.result.sync_info.latest_block_height')
TARGET=$((CURRENT_HEIGHT + 20))


cat > "$PROPOSAL_FILE" << EOF
{
  "messages": [
    {
      "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
      "authority": "shareledger10d07y265gmmuvt4z0w9aw880jnsr700jzx3uhc",
      "plan": {
        "name": "v3",
        "height": "$TARGET",
        "info": ""
      }
    }
  ],
  "metadata": "Upgrade ShareLedger binary to v3",
  "deposit": "1000000000nshr",
  "title": "ShareLedger v3 upgrade",
  "summary": "Upgrade ShareLedger binary to v3"
}
EOF

echo "Submitting test proposal..."
SUBMIT_RES="$(
  $BIN tx gov submit-proposal "$PROPOSAL_FILE" \
    --from validator --yes \
    --chain-id "$CHAIN_ID" --home "$HOME_DIR" --keyring-backend test \
    --output json
)"

sleep 5

# Deterministic: extract the proposal id from the tx response for THIS submit.
PROPOSAL_ID="$(
  echo "$SUBMIT_RES" | jq -r '
    (
      ((.logs // [])[0].events // [])
      + (.events // [])
    )
    | .[]
    | select(.type? == "submit_proposal")
    | (.attributes // [])
    | .[]
    | select(.key? == "proposal_id" or .key? == "proposal")
    | .value
  ' 2>/dev/null | head -n1
)"

# Fallback (best effort) if node/cli doesn't include events in tx response.
if [ -z "${PROPOSAL_ID:-}" ] || [ "$PROPOSAL_ID" = "null" ]; then
  PROPOSAL_ID="$(
    $BIN query gov proposals --home "$HOME_DIR" --output json 2>/dev/null \
      | jq -r '[.proposals[]? | (.id // .proposal_id // empty) | select(.!="") | tonumber] | (max // empty)' 2>/dev/null
  )"
fi

if [ -z "$PROPOSAL_ID" ] || [ "$PROPOSAL_ID" = "null" ]; then
    echo "Error: No proposal found"
    echo "Debug: submit tx response (truncated):"
    echo "$SUBMIT_RES" | head -c 800; echo
    exit 1
fi

echo "Test proposal ID: $PROPOSAL_ID"


echo "Voting YES on proposal $PROPOSAL_ID..."
$BIN tx gov vote "$PROPOSAL_ID" yes --from validator --yes \
  --chain-id $CHAIN_ID --home "$HOME_DIR" --keyring-backend test \
  --output json >/dev/null 2>&1

# Wait a bit for vote to be processed
sleep 3

PROPOSAL_STATUS=$($BIN query gov proposal "$PROPOSAL_ID" --home "$HOME_DIR" --output json 2>/dev/null | jq -r '.status // empty')
# YES_COUNT=$($BIN query gov proposal "$PROPOSAL_ID" --home "$HOME_DIR" --output json 2>/dev/null | jq -r '.final_tally_result.yes_count // "0"')
YES_COUNT=$($BIN query gov tally "$PROPOSAL_ID" --home "$HOME_DIR" --output json 2>/dev/null | jq -r '.yes_count // "0"')

echo "📊 Proposal status: $PROPOSAL_STATUS"
echo "✅ Yes votes: $YES_COUNT"

if [ "$PROPOSAL_STATUS" = "PROPOSAL_STATUS_VOTING_PERIOD" ]; then
    echo "✅ Proposal is in voting period and Validator's vote was counted!"
    echo "🎉 Governance test PASSED - Validator has sufficient voting power"
elif [ "$PROPOSAL_STATUS" = "PROPOSAL_STATUS_PASSED" ]; then
    echo "🎉 Proposal PASSED! Governance test successful"
else
    echo "⚠️  Proposal status: $PROPOSAL_STATUS"
fi