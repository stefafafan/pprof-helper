package main

import (
	"os"

	"golang.org/x/mod/modfile"
)

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
