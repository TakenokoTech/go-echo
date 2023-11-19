package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"sort"
	"strings"
	"time"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: strings.Join([]string{
			"time:${time_unix}",
			"method=${method}",
			"uri=${uri}",
			"status=${status}",
			"latency:${latency_human}",
			"error:${error}",
		}, ", ") + "\n",
	})
}

func RecoverMiddleware() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:           1 << 10, // 1 KB
		DisableErrorHandler: true,
		DisableStackAll:     true,
		DisablePrintStack:   true,
		LogLevel:            log.ERROR,
		LogErrorFunc:        nil,
	})
}

func RoutingList(e *echo.Echo) {
	time.Sleep(1000)
	fmt.Printf("==========================================\n")
	fmt.Println("[Route]")
	routes := e.Routes()
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Method < routes[j].Method
	})
	for _, route := range routes {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}
	fmt.Printf("==========================================\n")
}
