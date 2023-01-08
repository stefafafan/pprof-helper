package main

import "fmt"

func goji() {
	fmt.Println(`For goji, update the code like this:
import (
	"net/http/pprof"
)

func main() {
	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/debug/pprof/"), pprof.Index)
	mux.HandleFunc(pat.Get("/debug/pprof/cmdline"), pprof.Cmdline)
	mux.HandleFunc(pat.Get("/debug/pprof/profile"), pprof.Profile)
	mux.HandleFunc(pat.Get("/debug/pprof/symbol"), pprof.Symbol)
}`)
}
