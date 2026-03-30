//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Lint mg.Namespace

// All Run all linters
func (l Lint) All() error {
	mg.Deps(l.Go, l.Extra)
	return nil
}

// Extra lits everything that's not code
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
	return runMagefilesTool("gofumpt", "-l", "-w", ".")
}

// Golangcilint runs golangci-lint
func (Lint) Golangcilint() error {
	fmt.Println("running golangci-lint")
	return runMagefilesTool("golangci-lint", "run", "--fix")
}

// Vulncheck runs vulncheck
func (Lint) Vulncheck() error {
	fmt.Println("running vulncheck")
	return runMagefilesTool("govulncheck", "-show", "verbose", "./...")
}

// runMagefilesTool resolves a tool binary from the magefiles module and runs
// it in the repo root directory.
func runMagefilesTool(name string, args ...string) error {
	out, err := exec.Command("go", "-C", "magefiles", "tool", "-n", name).Output()
	if err != nil {
		return fmt.Errorf("resolving tool %s: %w", name, err)
	}
	bin := strings.TrimSpace(string(out))
	return runDirV(".", bin, args...)
}
