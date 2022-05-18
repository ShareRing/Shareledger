#bin/bash

echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add acc3 --recover --hd-path "m/44'/60'/0'/0/0"

for i in {1..100}
do
   shareledger tx swap out 0x97B98D335c28F9aD9c123E344a78F00C84146431 erc20 10shr --from user --yes
done

for i in {10..100..5}
do
  eval "shareledger tx swap approve $(( i - 1 )),$(( i - 2 )),$(( i - 3 )),$(( i - 4 )),$(( i - 5 ))  acc3 erc20 --trace --from authority --gas 500000 --yes"
done