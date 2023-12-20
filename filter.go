package main

type DexPoolFilter struct {
    Pools map[string]DexPoolInfo
}

func (f *DexPoolFilter) byLiquidity(minLiquidity float64) *DexPoolFilter {
    filteredPools := make(map[string]DexPoolInfo)
    for pair, info := range f.Pools {
        if info.Liquidity >= minLiquidity {
            filteredPools[pair] = info
        }
    }
    f.Pools = filteredPools
    return f
}

func (f *DexPoolFilter) byGasCost(maxGasCost float64) *DexPoolFilter {
    filteredPools := make(map[string]DexPoolInfo)
    for pair, info := range f.Pools {
        if info.GasCost <= maxGasCost {
            filteredPools[pair] = info
        }
    }
    f.Pools = filteredPools
    return f
}

func (f *DexPoolFilter) bySlippage(maxSlippage float64) *DexPoolFilter {
    filteredPools := make(map[string]DexPoolInfo)
    for pair, info := range f.Pools {
        if info.Slippage <= maxSlippage {
            filteredPools[pair] = info
        }
    }
    f.Pools = filteredPools
    return f
}

func (f *DexPoolFilter) byVolume(minVolume uint64) *DexPoolFilter {
    filteredPools := make(map[string]DexPoolInfo)
    for pair, info := range f.Pools {
        if info.Volume >= minVolume {
            filteredPools[pair] = info
        }
    }
    f.Pools = filteredPools
    return f
}
