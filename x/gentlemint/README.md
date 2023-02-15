# Gentlemint

> Calculate fee base on transaction type. This module make tx fee fixed base on stable token shrp

## Concepts

### 1) Exchange rate

Use to convert **shr** to **shrp**

For example: rate = 200 --> 200shr = 1shrp

_1 shrp = $1_

### 2) Level fee

| level  | amount (shrp) |
| ------ | ------------- |
| high   | 0.05          |
| medium | 0.03          |
| low    | 0.02          |
| zero   | 0             |

---

## Configuration

Via transaction by authority

### 1) Update global minimum_gas_price

Via gov submit proposal

```sh
shareledger tx gov submit-legacy-proposal param-change proposal.json
```

**proposal.json**

```json
{
  "title": "Gentlemint Param Change",
  "description": "Update min global fees",
  "changes": [
    {
      "subspace": "gentlemint",
      "key": "MinimumGasPricesParam",
      "value": [{ "denom": "nshr", "amount": "20000" }]
    }
  ],
  "deposit": "1000000000nshr"
}
```
