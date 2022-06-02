mongo -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB("admin");
db.auth("root", "ShareRingiscaring!");

db = db.getSiblingDB('Relayer');
db.states.insert({
    "network": "bsc",
    "lastScannedEventBlockNumbers": {
        "0x04ebff566170c907b29Fc33E9Cf0691faB87a168": 19551956,
        "0x9C83317a358dFDbe1e9B376A0E1CAB179C2c38AF": 19753469
    },
    "lastScannedBatchID": 0
});
db.states.insert({
    "network": "eth",
    "lastScannedEventBlockNumbers": {
        "0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560": 12290455,
        "0x0362fBA2BdA9Bd05f53d6C8CC72257919899A9Ac": 12316698
    },
    "lastScannedBatchID": 0
});
db.addresses.insert({
  "shareledgerAddress": "shareledger1006gjsnd449qy9mhmat7xwzqday0d7vsl24ur6",
  "AccIndex": 2,
  "mnemonicHash": "swap_user",
  "network": "eth",
  "result": "0x479d23Eef3b1Bb7F0D8c383CF5DC62B1936C0E5b"
});
db.addresses.insert({
  "shareledgerAddress": "shareledger1006gjsnd449qy9mhmat7xwzqday0d7vsl24ur6",
  "AccIndex": 3,
  "mnemonicHash": "swap_user",
  "network": "bsc",
  "result": "0x479d23Eef3b1Bb7F0D8c383CF5DC62B1936C0E5b"
});
EOF