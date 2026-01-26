# How to contribute

## Communication

- Issues: [GitHub](https://github.com/authzed/authzed-go/issues)
- Email: [Google Groups](https://groups.google.com/g/authzed-oss)
- Discord: [Zanzibar Discord](https://discord.gg/jTysUaxXzM)

All communication must follow our [Code of Conduct].

[Code of Conduct]: CODE-OF-CONDUCT.md

## Creating issues

If any part of the project has a bug or documentation mistakes, please let us know by opening an issue.
All bugs and mistakes are considered very seriously, regardless of complexity.

Before creating an issue, please check that an issue reporting the same problem does not already exist.
To make the issue accurate and easy to understand, please try to create issues that are:

- Unique -- do not duplicate existing bug report.
  Deuplicate bug reports will be closed.
- Specific -- include as much details as possible: which version, what environment, what configuration, etc.
- Reproducible -- include the steps to reproduce the problem.
  Some issues might be hard to reproduce, so please do your best to include the steps that might lead to the problem.
- Isolated -- try to isolate and reproduce the bug with minimum dependencies.
  It would significantly slow down the speed to fix a bug if too many dependencies are involved in a bug report.
  Debugging external systems that rely on this project is out of scope, but guidance or help using the project itself is fine.
- Scoped -- one bug per report.
  Do not follow up with another bug inside one report.

It may be worthwhile to read [Elika Etemad’s article on filing good bug reports][filing-good-bugs] before creating a bug report.

Maintainers might ask for further information to resolve an issue.

[filing-good-bugs]: http://fantasai.inkedblade.net/style/talks/filing-good-bugs/

## Contribution flow

This is a rough outline of what a contributor's workflow looks like:

- Create an issue
- Fork the project
- Create a branch from where to base the contribution -- this is almost always `main`
- Push changes into a branch of your fork
- Submit a pull request
- Respond to feedback from project maintainers

Creating new issues is one of the best ways to contribute.
You have no obligation to offer a solution or code to fix an issue that you open.
If you do decide to try and contribute something, please submit an issue first so that a discussion can occur to avoid any wasted efforts.

## Legal requirements

In order to protect both you and ourselves, all commits will require an explicit sign-off that acknowledges the [DCO].

Sign-off commits end with the following line:

```
Signed-off-by: Random J Developer <random@developer.example.org>
```

This can be done by using the `--signoff` (or `-s` for short) git flag to append this automatically to your commit message.
If you have already authored a commit that is missing the signed-off, you can amend or rebase your commits and force push them to GitHub.

[DCO]: /DCO

## Common tasks

### Testing

In order to build and test the project, the [latest stable version of Go] and knowledge of a [working Go environment] are required.

[latest stable version of Go]: https://golang.org/dl
[working Go environment]: https://golang.org/doc/code.html

```sh
go test -v ./...
```

### Adding dependencies

This project does not use anything other than the standard [Go modules] toolchain for managing dependencies.

[Go modules]: https://golang.org/ref/mod

```sh
go get github.com/org/newdependency@version
```

Continuous integration enforces that `go mod tidy` has been run.

### Updating generated Protobuf code

All [Protobuf] code is managed using [buf].
The constants in [magefiles/gen.go][gen-file] specify the [Buf Registry ref] that will be generated.
You can regenerate the code using the `mage` commands:

[Protobuf]: https://developers.google.com/protocol-buffers/
[buf]: https://docs.buf.build/installation
[Buf Registry ref]: https://buf.build/authzed/api/history
[gen-file]: https://github.com/authzed/authzed-go/blob/main/magefiles/gen.go#L26

```sh
mage gen:all
# or just the proto
mage gen:proto
```

#### Building with locally modified protos

If you need to test changes to the proto definitions before they're published to the Buf Registry:

1. Clone the [authzed/api] repository locally:
   ```sh
   git clone https://github.com/authzed/api.git
   ```

2. Make your changes to the `.proto` files in the `authzed/api` repository.

3. In the `authzed-go` repository, modify `magefiles/gen.go` to point to your local proto directory instead of the Buf Registry:
   ```go
   // Replace this line:
   bufRef := BufRepository + ":" + BufTag
   
   // With a path to your local api directory:
   bufRef := "/path/to/your/local/api"
   ```

4. Run the generation command:
   ```sh
   mage gen:proto
   ```

5. Before committing, remember to revert the changes to `magefiles/gen.go` so it points back to the Buf Registry reference.

**Note:** Generated version files will still reference the `BufRepository` and `BufTag` constants from the original registry reference, not your local path. This is expected for local testing and won't affect functionality.

[authzed/api]: https://github.com/authzed/api
