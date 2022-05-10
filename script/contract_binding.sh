#!/opt/homebrew/bin/bash
#go install github.com/ethereum/go-ethereum/cmd/abigen
abiDir=/Users/hoai/project/sharering/swap-contract-evm/abi
genDir=/Users/hoai/project/sharering/shareledger/pkg/swap/abi
for file in "$abiDir"/*.json; do
  [ -f "$file" ] || continue
  fileNameEx="${file##*/}"
  fileName="${fileNameEx%.*}"
  pkg=${fileName,,}
  abiPath="$genDir/$pkg"
  mkdir -p "$abiPath"
  abigen --pkg "$pkg" --abi="$file" --out "$abiPath/$fileName.gen.go"
done