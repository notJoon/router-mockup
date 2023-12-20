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

func TestFindExchangeRoutes(t *testing.T) {
    testCases := []struct {
        fromToken string
        toToken   string
    }{
        {"FOO", "BAZ"},
        {"BAR", "QUX"},
        {"BAZ", "FOO"},
		{"FOO", "QQQ"},
    }

    for _, tc := range testCases {
        routes, err := findExchangeRoutes(SwapRequest{FromToken: tc.fromToken, ToToken: tc.toToken}, dexPools)
        if err != nil {
            t.Errorf("findExchangeRoutes returned an error for %s/%s: %v", tc.fromToken, tc.toToken, err)
            continue
        }

        for _, route := range routes {
            formattedRoute := formatRoute(route)
			fmt.Printf("path for %s/%s: %s\n", tc.fromToken, tc.toToken, formattedRoute)
        }
    }
}
