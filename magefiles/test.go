//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

type Test mg.Namespace

// All runs all test suites
func (t Test) All() error {
	mg.Deps(t.Integration, t.Unit)
	return nil
}

// Unit runs the unit tests
func (t Test) Unit() error {
	fmt.Println("running unit tests")
	return goTest("./...", "-race", "-count", "1", "-timeout", "20m")
}

// Integration runs the integration tests
func (Test) Integration() error {
	fmt.Println("running integration tests")
	return goTest("./...", "-v", "-tags", "integration", "-timeout", "1m")
}
