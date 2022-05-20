#bin/bash

echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add relayer --recover # --keyring-backend os
echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add eth_signer --recover --hd-path "m/44'/60'/0'/0/0" # --keyring-backend os

echo "uphold snack joy universe tip gate click onion clean pigeon great sponsor bag inject skull kind excuse damage adapt shield myth ladder snow funny" | shareledger keys add relayer --recover


# enroll relayer
shareledger tx electoral enroll_relayer shareledger16p5u57wpd30j4k2kpg2rpxg3ye4m77lmjug4fq --from authority

#enroll approver
shareledger tx electoral enroll_approver shareledger16p5u57wpd30j4k2kpg2rpxg3ye4m77lmjug4fq --from authority

# Load shr to relayer
shareledger tx load shareledger16p5u57wpd30j4k2kpg2rpxg3ye4m77lmjug4fq 1000000shr --from authority --yes

# Load shr to user
shareledger tx load shareledger1006gjsnd449qy9mhmat7xwzqday0d7vsl24ur6 1000000shr --from authority --yes


# feed data
for i in {1..100}
do
   shareledger tx swap out 0x97B98D335c28F9aD9c123E344a78F00C84146431 erc20 10shr --from user --yes
done

for i in {10..100..10}
do
  eval "shareledger tx swap approve $((i)),$(( i - 1 )),$(( i - 2 )),$(( i - 3 )),$(( i - 4 )),$(( i - 5 )),$(( i - 6 )),$(( i - 7 )),$(( i - 8 )),$(( i - 9 )) eth_signer erc20 --trace --from authority --gas 50000000 --yes"
done