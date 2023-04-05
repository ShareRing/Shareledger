---

sidebar_position: 1
---

# `x/distributionx`

## Abstract

`x/distributionx` is mimic the behavior of `x/distribution` to customize logic to distribute fee:

   `Native module transaction`               `CosmWasm module transaction`
                                         +----------------------------------+
+------------------------------+         |  12.5% to smart contract owner   |
|                              |         |--------------------------------- |
|                              |         |  12.5% share to contract builders|
|     50% to developer pool    |         |--------------------------------- |
|                              |         |     25% to developer pool        |
|                              |         |                                  |
|-------------------------------         |--------------------------------- |
|                              |         |                                  |
|                              |         |                                  |
|       50% to validator       |         |         50% to validator         |
|                              |         |                                  |
|                              |         |                                  |
+------------------------------+         +----------------------------------+

## Contents

## TODO:

    - add unittest
