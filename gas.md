# Gas prices for Shareledger V2

gas_prices = 30,000nshr, gas_adjustify = 1.3

+-------------------|-----------|------+
| Transaction Type  | Gas used  | SHR  |
+===================|===========|======+
| bank.send         | 87,490    | 3.4  |
+-------------------|-----------|------+
| ibc.clientClient  | 93,451    | 3.7  |
| ibc.updateClient  | 94,873    | 3.7  |
+-------------------|-----------|------+
| wasm.StoreCode    | 1,092,746 | 42.5 |
| wasm.instantite   | 157,737   | 6.2  |
| wasm.execContract | 144,527   | 5.6  |
+-------------------|-----------|------+

The contract I used for WASM is just for reference; real gas usage might vary
based on the contract implementation.

=> then we config low: 30,000;med: 40,000; high: 50,000! for gas_prices
