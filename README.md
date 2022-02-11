# Authzed Go Client

[![GoDoc](https://godoc.org/github.com/authzed/authzed-go?status.svg)](https://godoc.org/github.com/authzed/authzed-go)
[![Docs](https://img.shields.io/badge/docs-authzed.com-%234B4B6C "Authzed Documentation")](https://docs.authzed.com)
[![Build Status](https://github.com/authzed/authzed-go/workflows/build/badge.svg)](https://github.com/authzed/authzed-go/actions)
[![Discord Server](https://img.shields.io/discord/844600078504951838?color=7289da&logo=discord "Discord Server")](https://discord.gg/jTysUaxXzM)
[![Twitter](https://img.shields.io/twitter/follow/authzed?color=%23179CF0&logo=twitter&style=flat-square)](https://twitter.com/authzed)

This repository houses the official Go client library for Authzed and SpiceDB.

[SpiceDB] is a database system for managing security-critical permissions checking.

SpiceDB acts as a centralized service that stores authorization data.
Once stored, data can be performantly queried to answer questions such as "Does this user have access to this resource?" and "What are all the resources this user has access to?".

[Authzed] operates the globally available, serverless database platform for SpiceDB.

Supported client API versions:
- [v1](https://buf.build/authzed/api/docs/main/authzed.api.v1)
- [v1alpha1](https://buf.build/authzed/api/docs/main/authzed.api.v1alpha1)
- "v0" - deprecated

You can find more info about the API in the [Authzed Documentation API Reference] or the [Authzed API Buf Registry repository].

See [CONTRIBUTING.md] for instructions on how to contribute and perform common tasks like building the project and running tests.

[SpiceDB]: https://github.com/authzed/spicedb
[Authzed]: https://authzed.com
[Authzed Documentation API Reference]: https://docs.authzed.com/reference/api
[Authzed API Buf Registry repository]: https://buf.build/authzed/api
[CONTRIBUTING.md]: CONTRIBUTING.md

## Getting Started

We highly recommend following the **[Protecting Your First App]** guide to learn the latest best practice to integrate an application with Authzed.

[Protecting Your First App]: https://docs.authzed.com/guides/first-app

## Basic Usage

### Installation

If you're using a modern version of [Go], run the following commands to add dependencies to your project:

```sh
go get github.com/authzed/authzed-go
go get github.com/authzed/grpcutil
```

[grpcutil] is not _strictly_ required, but greatly reduces the boilerplate required to create a client in the general case.

[Go]: https://golang.org/dl/
[grpcutil]: https://github.com/authzed/grpcutil

### Initializing a client

The [`NewClient()`] constructor is the recommended method for creating a client.

Because this library is using [gRPC] under the hood, you are free to leverage the wealth of functionality provided via [DialOptions].

In order to successfully connect, you will have to provide a [Bearer Token] with your own API Token from the [Authzed dashboard] in place of `t_your_token_here_1234567deadbeef` in the following example:

[`NewClient()`]: https://pkg.go.dev/github.com/authzed/authzed-go/v1#NewClient
[Bearer Token]: https://datatracker.ietf.org/doc/html/rfc6750#section-2.1
[Authzed Dashboard]: https://app.authzed.com
[gRPC]: https://grpc.io
[DialOptions]: https://pkg.go.dev/google.golang.org/grpc?utm_source=godoc#DialOption

```go
import (
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

...

client, err := authzed.NewClient(
	"grpc.authzed.com:443",
	grpcutil.WithSystemCerts(grpcutil.VerifyCA),
	grpcutil.WithBearerToken("t_your_token_here_1234567deadbeef"),
)
if err != nil {
	log.Fatalf("unable to initialize client: %s", err)
}
```

### Performing an API call

Requests and response types are located in a package under `proto/` respective to their API version.

Because of the verbosity of these types, we recommend writing your own functions/methods to create these types from your existing application's models.

```go
import (
	"github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

...

emilia := &pb.SubjectReference{Object: &v1.ObjectReference{
	ObjectType: "blog/user",
	ObjectId:  "emilia",
}}

firstPost := &pb.ObjectReference{
	ObjectType: "blog/post",
	ObjectId: "1",
}

resp, err := client.CheckPermission(ctx, &pb.CheckPermissionRequest{
	Resource: firstPost,
	Permission: "read",
	Subject: emilia,
})
if err != nil {
    log.Fatalf("failed to check permission: %s", err)
}

if resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
	log.Println("allowed!")
}
```
