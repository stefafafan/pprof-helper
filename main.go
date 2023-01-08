package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	var goModFile string
	flag.StringVar(&goModFile, "go-mod-file", "go.mod", "specify the path of the go.mod file")
	flag.Parse()

	goModBytes, err := os.ReadFile(goModFile)
	if err != nil {
		return
	}
	f, err := modfile.Parse(goModFile, goModBytes, nil)
	if err != nil {
		panic(err)
	}
	for _, v := range f.Require {
		fmt.Println(v.Mod)
	}
}
