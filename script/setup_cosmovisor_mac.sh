#!/usr/bin/env bash
# macOS local Cosmovisor upgrade setup for ShareLedger
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

HOME_DIR="${HOME_DIR:-$HOME/.shareledger}"
CHAIN_ID="${CHAIN_ID:-ShareRing-VoyagerNet}"

CURRENT_BIN="${CURRENT_BIN:-$HOME/go/bin/shareledger}"
NEW_BIN="${NEW_BIN:-$REPO_ROOT/build/shareledger}"

UPGRADE_NAME="${UPGRADE_NAME:-v3}"

INIT_SCRIPT="${INIT_SCRIPT:-$REPO_ROOT/script/init_dev_mac.sh}"

COSMOVISOR_BIN="${COSMOVISOR_BIN:-cosmovisor}"

DAEMON_NAME_DEFAULT="shareledger"
# IMPORTANT:
# - If your shell already has DAEMON_NAME set (e.g. from another chain), we *ignore* it by default.
# - Override explicitly via DAEMON_NAME_OVERRIDE=... (recommended).
# - Or keep inherited DAEMON_NAME by setting ALLOW_INHERIT_DAEMON_NAME=1.
if [ -n "${DAEMON_NAME_OVERRIDE:-}" ]; then
  DAEMON_NAME="$DAEMON_NAME_OVERRIDE"
elif [ -n "${DAEMON_NAME:-}" ] && [ "$DAEMON_NAME" != "$DAEMON_NAME_DEFAULT" ] && [ "${ALLOW_INHERIT_DAEMON_NAME:-0}" != "1" ]; then
  echo "Warning: DAEMON_NAME is already set to '$DAEMON_NAME' in your shell; ignoring it for this script."
  echo "- To override explicitly: DAEMON_NAME_OVERRIDE=shareledger"
  echo "- To keep your current value: ALLOW_INHERIT_DAEMON_NAME=1"
  DAEMON_NAME="$DAEMON_NAME_DEFAULT"
else
  DAEMON_NAME="${DAEMON_NAME:-$DAEMON_NAME_DEFAULT}"
fi

COSMOVISOR_HOME="$HOME_DIR/cosmovisor"
GENESIS_BIN_DIR="$COSMOVISOR_HOME/genesis/bin"
UPGRADE_BIN_DIR="$COSMOVISOR_HOME/upgrades/$UPGRADE_NAME/bin"

print_env () {
  cat <<EOF
export DAEMON_HOME="$HOME_DIR"
export DAEMON_NAME="$DAEMON_NAME"
export DAEMON_ALLOW_DOWNLOAD_BINARIES="${DAEMON_ALLOW_DOWNLOAD_BINARIES:-false}"
export DAEMON_RESTART_AFTER_UPGRADE="${DAEMON_RESTART_AFTER_UPGRADE:-true}"
export DAEMON_LOG_BUFFER_SIZE="${DAEMON_LOG_BUFFER_SIZE:-512}"
EOF
}

usage () {
  cat <<EOF
Usage:
  $0 [--print-env]

Examples:
  # Setup cosmovisor directories/binaries:
  $0

  # Apply env vars to your current shell:
  eval "\$($0 --print-env)"
  # or: $0 (writes env file) && source "$HOME_DIR/cosmovisor.env"

Environment overrides:
  HOME_DIR=...               (default: $HOME/.shareledger)
  UPGRADE_NAME=...           (default: v3)
  DAEMON_NAME_OVERRIDE=...   (recommended override)
  ALLOW_INHERIT_DAEMON_NAME=1 (use existing DAEMON_NAME if set)
EOF
}

case "${1:-}" in
  --print-env)
    print_env
    exit 0
    ;;
  -h|--help)
    usage
    exit 0
    ;;
  "")
    ;;
  *)
    echo "Unknown arg: $1" >&2
    usage >&2
    exit 2
    ;;
esac

echo "Effective config:"
echo "- HOME_DIR:      $HOME_DIR"
echo "- CHAIN_ID:      $CHAIN_ID"
echo "- DAEMON_NAME:   $DAEMON_NAME"
echo "- UPGRADE_NAME:  $UPGRADE_NAME"
echo "- CURRENT_BIN:   $CURRENT_BIN"
echo "- NEW_BIN:       $NEW_BIN"

echo "Checking binaries exist..."
test -f "$CURRENT_BIN" || { echo "Current binary not found: $CURRENT_BIN"; exit 1; }
test -f "$NEW_BIN" || { echo "New binary not found: $NEW_BIN"; exit 1; }

echo "Creating Cosmovisor directories..."
mkdir -p "$GENESIS_BIN_DIR"
mkdir -p "$UPGRADE_BIN_DIR"

echo "Installing genesis binary into Cosmovisor..."
cp -f "$CURRENT_BIN" "$GENESIS_BIN_DIR/$DAEMON_NAME"
chmod +x "$GENESIS_BIN_DIR/$DAEMON_NAME"

echo "Installing upgrade binary into Cosmovisor for upgrade name: $UPGRADE_NAME"
cp -f "$NEW_BIN" "$UPGRADE_BIN_DIR/$DAEMON_NAME"
chmod +x "$UPGRADE_BIN_DIR/$DAEMON_NAME"

echo "Optional: writing upgrade-info.json (some tooling expects it)"
cat > "$COSMOVISOR_HOME/upgrades/$UPGRADE_NAME/upgrade-info.json" <<EOF
{
  "name": "$UPGRADE_NAME",
  "info": "local cosmovisor test upgrade"
}
EOF

if [ "${RESET_HOME:-0}" = "1" ]; then
  echo "RESET_HOME=1 set. Removing $HOME_DIR"
  rm -rf "$HOME_DIR"
fi

if [ ! -d "$HOME_DIR/config" ] || [ ! -f "$HOME_DIR/config/genesis.json" ]; then
  echo "No existing chain home found. Running init script: $INIT_SCRIPT"
  bash "$INIT_SCRIPT"
else
  echo "Chain home already initialized at: $HOME_DIR"
fi

# Write an env file so you can 'source' it (since ./script/... can't modify your current shell environment).
ENV_FILE="$HOME_DIR/cosmovisor.env"
mkdir -p "$HOME_DIR"
print_env > "$ENV_FILE"

echo ""
echo "Cosmovisor env written to: $ENV_FILE"
echo "Apply it to your current shell with:"
echo "  source \"$ENV_FILE\""
echo "Or:"
echo "  eval \"\$($0 --print-env)\""

echo ""
echo "Note: running '$0' cannot persist exports into your current shell; use one of the commands above."
