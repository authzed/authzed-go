//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

type Test mg.Namespace

// All runs all test suites
func (t Test) All() error {
	mg.Deps(t.Integration)
	return nil
}

// Integration runs the unit tests
func (Test) Integration() error {
	fmt.Println("running unit tests")
	return goTest("./...", "-tags", "integration", "-timeout", "10m")
}
