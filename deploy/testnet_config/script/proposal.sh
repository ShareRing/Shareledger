shareledger tx gov submit-proposal software-upgrade v1.3.0-shareledger --upgrade-height 3229900 --description a --from node0 --title a -b block --deposit 1nshr -y
shareledger tx gov vote 1 yes -b block -y --from node0
shareledger tx gov vote 1 yes -b block -y --from node1
shareledger tx gov vote 1 yes -b block -y --from node2
shareledger tx gov vote 1 yes -b block -y --from node3