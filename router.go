package main

import (
	"errors"
	"fmt"
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

func ParseSwapRequest(request string) (SwapRequest, error) {
    params := strings.Split(request, "&")
    swapReq, err := parseQuery(params)
    if err != nil {
        return swapReq, err
    }

    return swapReq, nil
}

func parseQuery(q []string) (swapReq SwapRequest, err error) {
    if len(q) == 0 {
        return swapReq, EmptyQueryString
    }

    for _, p := range q {
        keyValue := strings.Split(p, "=")
        if len(keyValue) != 2 {
            return swapReq, fmt.Errorf("invalid parameter: %s", p)
        }

        if keyValue[1] == "" {
            return swapReq, InvalidQueryString
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
    Path       []string
    PathInfo   []DexPoolInfo
}

func FindExchangeRoutes(req SwapRequest, dexPools DexPoolInfoMap) ([]Route, error) {
    var findRoutes func(currentToken string, path []string, pathInfo []DexPoolInfo, visited map[string]bool) []Route
    findRoutes = func(currentToken string, path []string, pathInfo []DexPoolInfo, visited map[string]bool) []Route {
        if visited[currentToken] {
            return nil
        }

        // cycle detection
        visited[currentToken] = true
        var routes []Route

        for pair, info := range dexPools {
            tokens := strings.Split(pair, "/")
            if tokens[0] != currentToken {
                continue
            }

            newPath := append([]string(nil), path...)
            newPath = append(newPath, tokens[1])
            newPathInfo := append([]DexPoolInfo(nil), pathInfo...)
            newPathInfo = append(newPathInfo, info)

            if tokens[1] == req.ToToken {
                routes = append(routes, Route{Path: newPath, PathInfo: newPathInfo})
            } else {
                routes = append(routes, findRoutes(tokens[1], newPath, newPathInfo, copyMap(visited))...)
            }
        }

        visited[currentToken] = false
        return routes
    }

    return findRoutes(req.FromToken, []string{req.FromToken}, []DexPoolInfo{}, make(map[string]bool)), nil
}

// copyMap returns a copy of the original map.
func copyMap(original map[string]bool) map[string]bool {
    copy := make(map[string]bool)
    for key, value := range original {
        copy[key] = value
    }
    return copy
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
