package main

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
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
    var possibleRoutes []Route
    for fromToken := range dexPools {
        for toToken := range dexPools {
            if fromToken[:3] == req.FromToken && toToken[4:] == req.ToToken {
                intermediateToken := fromToken[4:]
                if intermediateToken == toToken[:3] {
                    possibleRoutes = append(possibleRoutes, Route{
                        Path:      []string{req.FromToken, intermediateToken, req.ToToken},
                        Liquidity: dexPools[fromToken] * dexPools[toToken],
                    })
                }
            }
        }

        directRouteKey := fmt.Sprintf("%s/%s", req.FromToken, req.ToToken)
        if liquidity, exists := dexPools[directRouteKey]; exists {
            possibleRoutes = append(possibleRoutes, Route{
                Path:      []string{req.FromToken, req.ToToken},
                Liquidity: liquidity,
            })
        }
    }

    var routes []Route
    for _, route := range possibleRoutes {
        if len(route.Path) <= 4 && route.Liquidity > 0 {
            routes = append(routes, route)
        }
    }

	// for debug
	for _, route := range routes {
		fmt.Println(formatRoute(route, req.FromToken, req.ToToken))
	}

    if len(routes) == 0 {
        return nil, fmt.Errorf("no available routes found")
    }

    return routes, nil
}

func formatRoute(route Route, fromToken string, toToken string) string {
    if len(route.Path) == 0 {
        return ""
    }

    if route.Path[0] != fromToken || route.Path[len(route.Path)-1] != toToken {
        return fmt.Sprintf("Invalid route: does not start with %s or end with %s", fromToken, toToken)
    }

    var pathStr strings.Builder
    pathStr.WriteString(fromToken)

    for i := 1; i < len(route.Path)-1; i += 2 {
        pathStr.WriteString(" -> ")
        pathStr.WriteString(route.Path[i])
        pathStr.WriteString("/")
        pathStr.WriteString(route.Path[i+1])
    }

    pathStr.WriteString(" -> ")
    pathStr.WriteString(toToken)

    return pathStr.String()
}