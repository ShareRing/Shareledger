mongo -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB("admin");
db.auth("root", "123");

db = db.getSiblingDB('Relayer');
db.states.insert({
    "network": "bsc",
    "lastScannedEventBlockNumbers": {
        "0x04ebff566170c907b29Fc33E9Cf0691faB87a168": 19551956
    },
    "lastScannedBatchID": 0
});
db.states.insert({
    "network": "eth",
    "lastScannedEventBlockNumbers": {
        "0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560": 12290455
    },
    "lastScannedBatchID": 0
});
EOF