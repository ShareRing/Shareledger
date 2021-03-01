BINARY=/shareledger/${BINARY:-slcli}
NODE=${NODE:-http://192.168.10.2:80}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'slcli' E.g.: -e BINARY=slcli_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

echo "Connecting to node ${NODE}"
"$BINARY" "$@" --node "$NODE" --trust-node --laddr tcp://0.0.0.0:1317

chmod 777 -R /shareledger

