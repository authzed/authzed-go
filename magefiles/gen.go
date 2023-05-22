//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Gen mg.Namespace

// All runs all generators in parallel
func (g Gen) All() error {
	mg.Deps(g.Proto)
	return nil
}

// Proto runs proto codegen
func (Gen) Proto() error {
	fmt.Println("generating buf")
	return sh.RunV("./buf.gen.yaml")
}
