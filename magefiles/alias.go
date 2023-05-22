//go:build mage

package main

var Aliases = map[string]interface{}{
	"test":     Test.Integration,
	"generate": Gen.All,
	"lint":     Lint.All,
	"gen":      Gen.All,
	"tidy":     Deps.Tidy,
}
