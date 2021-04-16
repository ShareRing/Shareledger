# Rate
Rate is the ration between SHRP and SHR.
eg:
rate = 200 --> 1SHRP = 200SHR

The rate decimal is 10^6, it means we add six 0 number after the rate do present the decimal.
Eg:
1SHRP=200SHR --> rate = 200*10^6 = 200,000,000
1SHRP=0.002SHR --> rate = 0.002*10^6 = 2,000

# Command
## Set rate
```
./slcli tx gentlemint set-exchange 1000000 --key-seed=./treasurer-seed.json  -y -b=block
```

## Get rate
```
./slcli query gentlemint get-exchange
```