package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

const (
	ECHO_V4 = "echo_v4"
	ECHO    = "echo"
	GOJI    = "goji"
	GIN     = "gin"
	BEEGO   = "beego"
	GOCHI   = "gochi"
)

var frameworks = map[string]string{
	"github.com/labstack/echo/v4": ECHO_V4,
	"github.com/labstack/echo":    ECHO,
	"goji.io":                     GOJI,
	"github.com/gin-gonic/gin":    GIN,
	"github.com/beego/beego/v2":   BEEGO,
	"github.com/go-chi/chi/v5":    GOCHI,
}

func main() {
	var goModFile string
	flag.StringVar(&goModFile, "go-mod-file", "go.mod", "specify the path of the go.mod file")
	flag.Parse()

	require, err := parsePackages(goModFile)
	if err != nil {
		return
	}

	framework := getFramework(require)
	if framework != "" {
		fmt.Println(framework)
	}
}

func parsePackages(path string) ([]*modfile.Require, error) {
	goModBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := modfile.Parse(path, goModBytes, nil)
	if err != nil {
		return nil, err
	}
	return f.Require, nil
}

func getFramework(requires []*modfile.Require) string {
	for _, req := range requires {
		foo := frameworks[req.Mod.Path]
		if foo != "" {
			return foo
		}
	}
	return ""
}
