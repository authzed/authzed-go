//go:build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

// codegenModules lists the modules that provide both a codegen plugin (run
// from the magefiles module via `go run`; see magefiles/buf.gen.yaml) and the
// runtime library that the generated code links against (a dependency of the
// root module). Their versions must match so that generated code and runtime
// agree; keep this list in sync with magefiles/buf.gen.yaml.
//
// protoc-gen-go-grpc is intentionally absent: it lives in its own module
// (google.golang.org/grpc/cmd/protoc-gen-go-grpc) with a version series
// unrelated to the google.golang.org/grpc runtime, so there is nothing to
// compare.
var codegenModules = []string{
	"google.golang.org/protobuf",
	"github.com/grpc-ecosystem/grpc-gateway/v2",
	"github.com/planetscale/vtprotobuf",
	"github.com/envoyproxy/protoc-gen-validate",
}

// Depsync checks that codegen plugin versions in magefiles/go.mod match go.mod
func (Lint) Depsync() error {
	fmt.Println("checking codegen plugin versions against go.mod")

	rootVersions, err := moduleVersions("go.mod")
	if err != nil {
		return err
	}
	mageVersions, err := moduleVersions("magefiles/go.mod")
	if err != nil {
		return err
	}

	var mismatches []string
	for _, mod := range codegenModules {
		rootVersion, inRoot := rootVersions[mod]
		mageVersion, inMage := mageVersions[mod]
		if !inRoot || !inMage {
			// The module is no longer used by one of the go.mod files;
			// this list and magefiles/buf.gen.yaml likely need updating.
			continue
		}
		if rootVersion != mageVersion {
			mismatches = append(mismatches, fmt.Sprintf(
				"%s: go.mod has %s but magefiles/go.mod has %s",
				mod, rootVersion, mageVersion))
		}
	}

	if len(mismatches) > 0 {
		return fmt.Errorf(
			"codegen plugin versions are out of sync with go.mod:\n\t%s\n"+
				"align them (e.g. `go get -C magefiles <module>@<version>`), then run `mage deps:tidy` and `mage gen:proto`",
			strings.Join(mismatches, "\n\t"))
	}
	return nil
}

func moduleVersions(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", path, err)
	}
	file, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", path, err)
	}

	versions := make(map[string]string, len(file.Require))
	for _, req := range file.Require {
		versions[req.Mod.Path] = req.Mod.Version
	}
	return versions, nil
}
