#bin/bash
#
echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add relayer --recover # --keyring-backend os
echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add eth_signer --recover --hd-path "m/44'/60'/0'/0/0" # --keyring-backend os

echo "logic fade bike misery female father false speak code immune improve key food enter night timber kick spare amused miss expire bottom walk century" | shareledger keys add swap_manager --recover
#
echo "uphold snack joy universe tip gate click onion clean pigeon great sponsor bag inject skull kind excuse damage adapt shield myth ladder snow funny" | shareledger keys add swap_user --recover
#
#
# enroll relayer
shareledger tx electoral enroll-relayer $(shareledger keys show relayer -a) --from authority --yes

#enroll approver
shareledger tx electoral enroll-approver $(shareledger keys show relayer -a) --from authority --yes

#enroll swap_manager
shareledger tx electoral enroll-swap-managers $(shareledger keys show swap_manager -a) --from authority --yes

# Load shr to relayer
shareledger tx gm load $(shareledger keys show relayer -a) 1000000shr --from authority --yes

# Load shr to user
shareledger tx gm load $(shareledger keys show swap_user -a) 10000000shr --from authority --yes

# Add Schema
shareledger tx swap schema eth '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}' 2shr 2shr 2 --from authority --yes
shareledger tx swap schema bsc '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x61","verifyingContract":"0x04ebff566170c907b29Fc33E9Cf0691faB87a168","salt":""}}' 2shr 2shr 18 --from authority --yes


# feed data
#for i in {941..1100}
#do
#   shareledger tx swap out 0x97B98D335c28F9aD9c123E344a78F00C84146431 bsc 0.$((i))shr --from swap_user --yes
#done
#
#for i in {950..1100..5}
#do
#  eval "shareledger tx swap approve $((i)),$((i-4)),$((i-3)),$((i-2)),$((i-1)) eth_signer bsc --trace --from authority --gas 50000000 --yes"
#done