# Gentlemint

> Calculate fee base on transaction type. This module make tx fee fixed base on stable token shrp

## Concepts

### Exchange rate

Use to convert **shr** to **shrp**

For example: rate = 200 --> 200shr = 1shrp

_1 shrp = $1_

### Level fee

| level  | amount (shrp) |
| ------ | ------------- |
| high   | 0.05          |
| medium | 0.03          |
| low    | 0.02          |
| zero   | 0             |

---

## Configuration

Via transaction by authority
