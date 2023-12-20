package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	EmptyQueryString error	 = errors.New("query string is empty")
	InvalidQueryString		 = errors.New("invalid query string")
	FromIsNotValid			 = errors.New("invalid from")
	ToIsNotValid			 = errors.New("invalid to")
	AmountIsNotValid		 = errors.New("invalid amount")
)

type SwapRequest struct {
    FromToken  string
    ToToken    string
    Amount     float64
    Slippage   float64
    Trader     string
    MinReceive float64
    GasPrice   float64
}

func parseSwapRequest(request string) (SwapRequest, error) {
    var swapReq SwapRequest
    params := strings.Split(request, "&")

    for _, p := range params {
        keyValue := strings.Split(p, "=")
        if len(keyValue) != 2 {
            return swapReq, fmt.Errorf("invalid parameter: %s", p)
        }

        key, value := keyValue[0], keyValue[1]
        switch key {
        case "from":
            swapReq.FromToken = value
        case "to":
            swapReq.ToToken = value
        case "amount":
            if val, err := strconv.ParseFloat(value, 64); err == nil {
                swapReq.Amount = val
            } else {
                return swapReq, err
            }
        case "slippage":
            if val, err := strconv.ParseFloat(value, 64); err == nil {
                swapReq.Slippage = val
            } else {
                return swapReq, err
            }
        case "trader":
            swapReq.Trader = value
        case "minReceive":
            if val, err := strconv.ParseFloat(value, 64); err == nil {
                swapReq.MinReceive = val
            } else {
                return swapReq, err
            }
        case "gasPrice":
            if val, err := strconv.ParseFloat(value, 64); err == nil {
                swapReq.GasPrice = val
            } else {
                return swapReq, err
            }
        default:
            return swapReq, fmt.Errorf("unknown parameter: %s", key)
        }
    }

    return swapReq, nil
}

type Route struct {
    Path      []string
    Liquidity float64
}

func findExchangeRoutes(req SwapRequest, dexPools map[string]float64) ([]Route, error) {
    var findRoutes func(currentToken string, path []string, liquidity float64) []Route
    findRoutes = func(currentToken string, path []string, liquidity float64) []Route {
        if len(path) > 3 {
            return nil
        }

        var routes []Route
        for pair, poolLiquidity := range dexPools {
            tokens := strings.Split(pair, "/")
            if tokens[0] != currentToken {
                continue
            }

            newPath := append([]string(nil), path...)
            newPath = append(newPath, pair)
            newLiquidity := math.Min(liquidity, poolLiquidity)

            if tokens[1] == req.ToToken {
                routes = append(routes, Route{Path: newPath, Liquidity: newLiquidity})
            } else {
                routes = append(routes, findRoutes(tokens[1], newPath, newLiquidity)...)
            }
        }
        return routes
    }

    return findRoutes(req.FromToken, []string{req.FromToken}, math.MaxFloat64), nil
}

func formatRoute(route Route) string {
    if len(route.Path) == 0 {
        return ""
    }

    var pathStr strings.Builder
    for i, pair := range route.Path {
        if i > 0 {
            pathStr.WriteString(" -> ")
        }
        pathStr.WriteString(pair)

        if i == len(route.Path)-1 {
            tokens := strings.Split(pair, "/")
            if len(tokens) == 2 {
                pathStr.WriteString(" -> ")
                pathStr.WriteString(tokens[1])
            }
        }
    }

    return pathStr.String()
}
