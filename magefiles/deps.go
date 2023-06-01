//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

var goModules = []string{".", "magefiles"}

type Deps mg.Namespace

// Tidy go mod tidy all go modules
func (Deps) Tidy() error {
	for _, mod := range goModules {
		fmt.Println("tidying", mod)
		if err := runDirV(mod, "go", "mod", "tidy"); err != nil {
			return err
		}
	}

	return nil
}

// Update go get -u all go dependencies
func (Deps) Update() error {
	for _, mod := range goModules {
		fmt.Println("updating", mod)
		if err := runDirV(mod, "go", "get", "-u", "-t", "-tags", "tools", "./..."); err != nil {
			return err
		}

		if err := runDirV(mod, "go", "mod", "tidy"); err != nil {
			return err
		}
	}

	return nil
}
