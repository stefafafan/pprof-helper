package main

import (
	"flag"
	"fmt"
)

const (
	ECHO  = "echo"
	GOJI  = "goji"
	GIN   = "gin"
	BEEGO = "beego"
	GOCHI = "gochi"
)

var frameworks = map[string]string{
	"github.com/labstack/echo/v4": ECHO,
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
	switch framework {
	case ECHO:
		echo()
	case GOJI:
	case BEEGO:
	case GOCHI:
	default:
		fmt.Printf("framework not supported: %s\n", framework)
	}
}
