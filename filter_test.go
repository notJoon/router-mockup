package main

import (
	"testing"
)

func TestDexPoolFilter_FilterByLiquidity(t *testing.T) {
	filter := &DexPoolFilter{
		Pools: map[string]DexPoolInfo{
			"pair1": { Liquidity: 10.0 },
			"pair2": { Liquidity: 5.0 },
			"pair3": { Liquidity: 15.0 },
		},
	}

	minLiquidity := 10.0
	filtered := filter.byLiquidity(minLiquidity)

	expected := map[string]DexPoolInfo{
		"pair1": { Liquidity: 10.0 },
		"pair3": { Liquidity: 15.0 },
	}

	if len(filtered.Pools) != len(expected) {
		t.Errorf("Expected %d filtered pools, but got %d", len(expected), len(filtered.Pools))
	}

	for pair, info := range expected {
		if filtered.Pools[pair] != info {
			t.Errorf("Expected filtered pool %s to be %+v, but got %+v", pair, info, filtered.Pools[pair])
		}
	}
}

func TestDexPoolFilter_FilterByGasCost(t *testing.T) {
	filter := &DexPoolFilter{
		Pools: map[string]DexPoolInfo{
			"pair1": {GasCost: 0.5},
			"pair2": {GasCost: 0.8},
			"pair3": {GasCost: 0.3},
		},
	}

	maxGasCost := 0.6
	filtered := filter.byGasCost(maxGasCost)

	expected := map[string]DexPoolInfo{
		"pair1": {GasCost: 0.5},
		"pair3": {GasCost: 0.3},
	}

	if len(filtered.Pools) != len(expected) {
		t.Errorf("Expected %d filtered pools, but got %d", len(expected), len(filtered.Pools))
	}

	for pair, info := range expected {
		if filtered.Pools[pair] != info {
			t.Errorf("Expected filtered pool %s to be %+v, but got %+v", pair, info, filtered.Pools[pair])
		}
	}
}

func TestDexPoolFilter_FilterBySlippage(t *testing.T) {
	filter := &DexPoolFilter{
		Pools: map[string]DexPoolInfo{
			"pair1": {Slippage: 0.01},
			"pair2": {Slippage: 0.02},
			"pair3": {Slippage: 0.03},
		},
	}

	maxSlippage := 0.02
	filtered := filter.bySlippage(maxSlippage)

	expected := map[string]DexPoolInfo{
		"pair1": {Slippage: 0.01},
		"pair2": {Slippage: 0.02},
	}

	if len(filtered.Pools) != len(expected) {
		t.Errorf("Expected %d filtered pools, but got %d", len(expected), len(filtered.Pools))
	}

	for pair, info := range expected {
		if filtered.Pools[pair] != info {
			t.Errorf("Expected filtered pool %s to be %+v, but got %+v", pair, info, filtered.Pools[pair])
		}
	}
}
