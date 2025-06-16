//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Lint mg.Namespace

// All Run all linters
func (l Lint) All() error {
	mg.Deps(l.Go, l.Extra)
	return nil
}

// Extra lints everything that's not code
func (l Lint) Extra() error {
	mg.Deps(l.Yaml)
	return nil
}

// Yaml lints yaml
func (Lint) Yaml() error {
	mg.Deps(checkDocker)
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return sh.RunV("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/src:ro", cwd),
		"cytopia/yamllint:1", "-c", "/src/.yamllint", "/src")
}

// Go runs all go linters
func (l Lint) Go() error {
	mg.Deps(l.Gofumpt, l.Golangcilint, l.Vulncheck)
	return nil
}

// Gofumpt runs gofumpt
func (Lint) Gofumpt() error {
	fmt.Println("formatting go")
	return runDirV(".", "go", "run", "mvdan.cc/gofumpt", "-l", "-w", "..")
}

// Golangcilint runs golangci-lint
func (Lint) Golangcilint() error {
	fmt.Println("running golangci-lint")
	return runDirV(".", "go", "run", "github.com/golangci/golangci-lint/v2/cmd/golangci-lint", "run", "--fix")
}

// Vulncheck runs vulncheck
func (Lint) Vulncheck() error {
	fmt.Println("running vulncheck")
	return runDirV(".", "go", "run", "golang.org/x/vuln/cmd/govulncheck", "-show", "verbose", "./...")
}
