shareledger tx electoral enroll-relayer $(shareledger keys show relayer -a) --from authority --yes -b block
shareledger tx electoral enroll-approver $(shareledger keys show relayer -a) --from authority --yes -b block
shareledger tx electoral enroll-swap-managers $(shareledger keys show swap_manager -a) --from authority --yes -b block
# Load shr to relayer
shareledger tx gm load $(shareledger keys show relayer -a) 1000000shr --from authority --yes -b block
# Load shr to user
shareledger tx gm load $(shareledger keys show swap_user -a) 10000000shr --from authority --yes -b block

# Add Schema
shareledger tx swap schema eth '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}' 2shr 2shr 2 --from authority --yes -b block
shareledger tx swap schema bsc '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x61","verifyingContract":"0x04ebff566170c907b29Fc33E9Cf0691faB87a168","salt":""}}' 2shr 2shr 18 --from authority --yes -b block