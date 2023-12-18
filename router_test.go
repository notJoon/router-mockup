package main

import (
	"fmt"
	"testing"
)

func TestParseSwapRequest(t *testing.T) {
    testRequest := "from=FOO&to=BAR&amount=100.0&slippage=0.01&trader=0xExampleAddress&minReceive=99.0&gasPrice=50"
    expected := SwapRequest{
        FromToken:  "FOO",
        ToToken:    "BAR",
        Amount:     100.0,
        Slippage:   0.01,
        Trader:     "0xExampleAddress",
        MinReceive: 99.0,
        GasPrice:   50.0,
    }

    result, err := parseSwapRequest(testRequest)
    if err != nil {
        t.Errorf("parseSwapRequest returned an error: %v", err)
    }
    if result != expected {
        t.Errorf("parseSwapRequest = %v; want %v", result, expected)
    }
}

func TestFindExchangeRoutes(t *testing.T) {
    var dexPools = map[string]float64 {
		"FOO/BAR": 100000.0,
		"BAR/FOO": 100000.0,
		"FOO/BAZ": 200000.0,
		"BAZ/FOO": 200000.0,
		"FOO/QUX": 300000.0,
		"QUX/FOO": 300000.0,
		"BAR/BAZ": 400000.0,
		"BAZ/BAR": 400000.0,
		"BAR/QUX": 500000.0,
		"QUX/BAR": 500000.0,
		"BAZ/QUX": 600000.0,
		"QUX/BAZ": 600000.0,
	}

    testCases := []struct {
        name         string
        request      SwapRequest
        expectedPath []string
        wantErr      bool
    }{
        {
			name: "FOO -> BAR",
			request: SwapRequest{
				FromToken:  "FOO",
				ToToken:    "BAR",
				Amount:     100.0,
				Slippage:   0.01,
				Trader:     "0xExampleAddress",
				MinReceive: 99.0,
				GasPrice:   50.0,
			},
			expectedPath: []string{"FOO", "BAR"},
			wantErr:      false,
		},
		{
			name: "FOO -> QUX",
			request: SwapRequest{
				FromToken:  "FOO",
				ToToken:    "QUX",
				Amount:     200.0,
				Slippage:   0.02,
				Trader:     "0xAnotherExampleAddress",
				MinReceive: 198.0,
				GasPrice:   60.0,
			},
			expectedPath: []string{"FOO", "QUX"},
			wantErr:      false,
		},
		{
			name: "BAR -> QUX",
			request: SwapRequest{
				FromToken:  "BAR",
				ToToken:    "QUX",
				Amount:     300.0,
				Slippage:   0.03,
				Trader:     "0xYetAnotherExampleAddress",
				MinReceive: 291.0,
				GasPrice:   70.0,
			},
			expectedPath: []string{"BAR", "QUX"},
			wantErr:      false,
		},
		{
			name: "QUX -> FOO",
			request: SwapRequest{
				FromToken:  "QUX",
				ToToken:    "FOO",
				Amount:     400.0,
				Slippage:   0.04,
				Trader:     "0xYetAnotherExampleAddress",
				MinReceive: 384.0,
				GasPrice:   80.0,
			},
			expectedPath: []string{"QUX", "FOO"},
			wantErr:      false,
		},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            route, err := findExchangeRoutes(tc.request, dexPools)
			fmt.Println("route: ", route)
            if (err != nil) != tc.wantErr {
                t.Errorf("findExchangeRoutes() error = %v, wantErr %v", err, tc.wantErr)
                return
            }
        })
    }
}
