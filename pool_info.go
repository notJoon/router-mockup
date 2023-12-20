package main

type DexPoolInfo struct {
    Liquidity float64
    GasCost   float64
    Slippage  float64
    Volume    uint64
    TokenA    string
    TokenB    string
    ReserveA  float64
    ReserveB  float64
}

type DexPoolInfoMap map[string]DexPoolInfo

var DexPools = map[string]DexPoolInfo{
    "FOO/BAR": {Liquidity: 100000.0, GasCost: 48.80, Slippage: 0.041, Volume: 100000, TokenA: "FOO", TokenB: "BAR", ReserveA: 9713.75, ReserveB: 9945.93},
    "BAR/FOO": {Liquidity: 100000.0, GasCost: 25.52, Slippage: 0.045, Volume: 100000, TokenA: "BAR", TokenB: "FOO", ReserveA: 5081.82, ReserveB: 7389.13},
    "FOO/BAZ": {Liquidity: 200000.0, GasCost: 36.04, Slippage: 0.027, Volume: 200000, TokenA: "FOO", TokenB: "BAZ", ReserveA: 9183.88, ReserveB: 6486.36},
    "BAZ/FOO": {Liquidity: 200000.0, GasCost: 13.79, Slippage: 0.036, Volume: 200000, TokenA: "BAZ", TokenB: "FOO", ReserveA: 2408.20, ReserveB: 9458.81},
    "BAR/BAZ": {Liquidity: 300000.0, GasCost: 26.11, Slippage: 0.046, Volume: 300000, TokenA: "BAR", TokenB: "BAZ", ReserveA: 9713.75, ReserveB: 9945.93},
    "BAZ/BAR": {Liquidity: 300000.0, GasCost: 42.00, Slippage: 0.027, Volume: 300000, TokenA: "BAZ", TokenB: "BAR", ReserveA: 5081.82, ReserveB: 7389.13},
    "BAR/QUX": {Liquidity: 400000.0, GasCost: 34.85, Slippage: 0.035, Volume: 400000, TokenA: "BAR", TokenB: "QUX", ReserveA: 9183.88, ReserveB: 6486.36},
    "QUX/BAR": {Liquidity: 400000.0, GasCost: 19.01, Slippage: 0.014, Volume: 400000, TokenA: "QUX", TokenB: "BAR", ReserveA: 2408.20, ReserveB: 9458.81},
    "BAZ/QUX": {Liquidity: 500000.0, GasCost: 49.17, Slippage: 0.023, Volume: 500000, TokenA: "BAZ", TokenB: "QUX", ReserveA: 9713.75, ReserveB: 9945.93},
    "QUX/BAZ": {Liquidity: 500000.0, GasCost: 24.02, Slippage: 0.033, Volume: 500000, TokenA: "QUX", TokenB: "BAZ", ReserveA: 5081.82, ReserveB: 7389.13},
    "FOO/QUX": {Liquidity: 600000.0, GasCost: 24.52, Slippage: 0.021, Volume: 600000, TokenA: "FOO", TokenB: "QUX", ReserveA: 9183.88, ReserveB: 6486.36},
    "QUX/FOO": {Liquidity: 600000.0, GasCost: 19.50, Slippage: 0.017, Volume: 600000, TokenA: "QUX", TokenB: "FOO", ReserveA: 2408.20, ReserveB: 9458.81},
}

