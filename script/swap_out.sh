bin/bash

#echo "glory fame cat rebuild spider wide forget easy bundle report lava comfort kiss cry tooth dwarf concert direct giggle scale caution cinnamon bundle display" | shareledger keys add khang --recover --hd-path "m/44'/60'/0'/0/0"
#
#for i in {4001..4050}
#do
#   shareledger tx swap out 0x97B98D335c28F9aD9c123E344a78F00C84146431 eth $(( $RANDOM % 20 + 1 ))shr --from user --yes
#done

for i in {4010..4050..10}
do
  eval "shareledger tx swap approve $((i)),$(( i - 1 )),$(( i - 2 )),$(( i - 3 )),$(( i - 4 )),$(( i - 5 )),$(( i - 6 )),$(( i - 7)),$(( i - 8)),$(( i - 9 )) khang eth --trace --from authority --gas 500000000 --yes"
done