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

    result, err := ParseSwapRequest(testRequest)
    if err != nil {
        t.Errorf("parseSwapRequest returned an error: %v", err)
    }
    if result != expected {
        t.Errorf("parseSwapRequest = %v; want %v", result, expected)
    }
}

func TestInvalidSwapRequestHasReceived(t *testing.T) {
    queries := []struct {
        name string
        req  string
    }{
        {name: "empty", req: ""},
        {name: "missing from value", req: "from=&to=BAR&amount=100.0&slippage=0.01&trader=0xExampleAddress&minReceive=99.0&gasPrice=50"},
        {name: "missing to value", req: "from=FOO&to=&amount=100.0&slippage=0.01&trader=0xExampleAddress&minReceive=99.0&gasPrice=50"},
        {name: "missing amount value", req: "from=FOO&to=BAR&amount=&slippage=0.01&trader=0xExampleAddress&minReceive=99.0&gasPrice=50"},
        {name: "missing slippage value", req: "from=FOO&to=BAR&amount=100.0&slippage=&trader=0xExampleAddress&minReceive=99.0&gasPrice=50"},
        {name: "missing trader value", req: "from=FOO&to=BAR&amount=100.0&slippage=0.01&trader=&minReceive=99.0&gasPrice=50"},
        {name: "missing minReceive value", req: "from=FOO&to=BAR&amount=100.0&slippage=0.01&trader=0xExampleAddress&minReceive=&gasPrice=50"},
        {name: "missing gasPrice value", req: "from=FOO&to=BAR&amount=100.0&slippage=0.01&trader=0xExampleAddress&minReceive=99.0&gasPrice="},
    }

    for _, tc := range queries {
        _, err := ParseSwapRequest(tc.req)
        if err == nil {
            t.Errorf("parseSwapRequest(%s) did not return an error", tc.name)
        }
    }
}

func TestFindExchangeRoutes(t *testing.T) {
    testRequest := SwapRequest{
        FromToken: "FOO",
        ToToken:   "BAZ",
        Amount:    1000.0,
    }

    testDexPools := DexPools

    routes, err := FindExchangeRoutes(testRequest, testDexPools)
    if err != nil {
        t.Errorf("FindExchangeRoutes returned an error: %v", err)
    }

    if len(routes) == 0 {
        t.Errorf("No routes found for request %+v", testRequest)
    }

    for _, route := range routes {
        if len(route.Path) == 0 || len(route.PathInfo) == 0 {
            t.Errorf("Route or PathInfo is empty for route: %+v", route)
        }

        if containsCycle(route.Path) {
            t.Errorf("Route contains a cycle: %+v", route)
        }

        format := formatRoute(route)
        fmt.Println(format)
    }
}

func containsCycle(path []string) bool {
    visited := make(map[string]bool)
    for _, token := range path {
        if _, found := visited[token]; found {
            return true
        }
        visited[token] = true
    }
    return false
}
