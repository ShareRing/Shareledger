shareledger config keyring-backend test
shareledger config chain-id ShareRing-Lifestyle

echo 'work like grass pyramid august topic exit wood reunion until retire frown bean cherry stage attack shed oxygen chronic return kiss hint hat future' | shareledger keys add authority --recover --index 0
echo 'bird pizza tobacco omit cricket noodle hold wagon opinion shiver scout nature discover almost permit ceiling endless total sight cattle crisp calm popular flat' | shareledger keys add treasurer --recover --index 0
echo 'decline electric decade treat scissors floor fade fade exile swim destroy unusual noodle cabin toy print cover limb old report that balance pelican member' | shareledger keys add operator --recover --index 0
echo 'mule inform panel original wagon guard open inquiry inner vivid latin live lumber solar tank claw celery tattoo focus raven bunker cargo turn noble' | shareledger keys add shrp-loader --recover --index 0

echo 'lawsuit icon drill call jaguar party text chunk woman mention dawn frost net illness tail garlic exotic orange cage grape analyst jealous road cupboard' | shareledger keys add node0 --recover --index 0
echo 'honey salad ride tackle border mail safe upset grape oven define coffee erase aunt stem skirt urban below ignore skill key indicate never rule' | shareledger keys add node1 --recover --index 0
echo 'avocado blue utility tell hawk gorilla auction morning dolphin offer beyond unhappy orphan soccer asset route kite edge question tuna fragile company total minute' | shareledger keys add node2 --recover --index 0
echo 'mango stuff casino shoulder tattoo labor civil master guitar blouse coral cabbage rose obscure winter copy learn hour explain surge tip disorder track life' | shareledger keys add node3 --recover --index 0

#swap features
echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add relayer --recover
echo "upper two old trip wolf member fruit end coyote stone gospel knee" | shareledger keys add eth_signer --recover --hd-path "m/44'/60'/0'/0/0"
echo "logic fade bike misery female father false speak code immune improve key food enter night timber kick spare amused miss expire bottom walk century" | shareledger keys add swap_manager --recover
echo "uphold snack joy universe tip gate click onion clean pigeon great sponsor bag inject skull kind excuse damage adapt shield myth ladder snow funny" | shareledger keys add swap_user --recover

shareledger tx electoral enroll-relayer $(shareledger keys show relayer -a) --from authority --yes
shareledger tx electoral enroll-approver $(shareledger keys show relayer -a) --from authority --yes
shareledger tx electoral enroll-swap-managers $(shareledger keys show swap_manager -a) --from authority --yes
# Load shr to relayer
shareledger tx gm load $(shareledger keys show relayer -a) 1000000shr --from authority --yes
# Load shr to user
shareledger tx gm load $(shareledger keys show swap_user -a) 10000000shr --from authority --yes

# Add Schema
shareledger tx swap schema eth '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}' 2shr 2shr 2 --from authority --yes
shareledger tx swap schema bsc '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x61","verifyingContract":"0x04ebff566170c907b29Fc33E9Cf0691faB87a168","salt":""}}' 2shr 2shr 18 --from authority --yes