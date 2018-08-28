#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
rootdir="$workspace/src/github.com/sharering"
proj="shareledger"
if [ ! -L "$rootdir/$proj" ]; then
    echo "Create link"
    mkdir -p "$rootdir"
    cd "$rootdir"
    ln -s ../../../../../. $proj
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$rootdir/$proj"
PWD="$rootdir/$proj"

# Launch the arguments with the configured environment.
exec "$@"
