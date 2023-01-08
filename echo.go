package main

import "fmt"

func echo() {
	fmt.Println(`For echo, first install the package (sevenNt/echo-pprof now supports echo v4).

go get -u github.com/sevenNt/echo-pprof

Then, update the code:

import (
	"github.com/sevenNt/echo-pprof"
)

func main() {
	e := echo.New()
	echopprof.Wrap(e)
}`)
}
