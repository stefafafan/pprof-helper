package main

import "fmt"

func echoV4() {
	fmt.Println(`For echo v4, add the following code:

import (
	"net/http/pprof"
)

func main() {
	e := echo.New()

	// ...

	pprofGroup := e.Group("/debug/pprof")
	pprofGroup.Any("/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	pprofGroup.Any("/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	pprofGroup.Any("/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	pprofGroup.Any("/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))
	pprofGroup.Any("/*", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
}`)

}
