# Authzed Go Client

[![GoDoc](https://godoc.org/github.com/authzed/authzed-go?status.svg)](https://godoc.org/github.com/authzed/authzed-go)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Build Status](https://github.com/authzed/authzed-go/workflows/build/badge.svg)](https://github.com/authzed/authzed-go/actions)
[![Discord Server](https://img.shields.io/discord/844600078504951838?color=7289da&logo=discord "Discord Server")](https://discord.gg/jTysUaxXzM)
[![Twitter](https://img.shields.io/twitter/follow/authzed?color=%23179CF0&logo=twitter&style=flat-square)](https://twitter.com/authzed)

This repository houses the Go client library for Authzed.

[Authzed] is a database and service that stores, computes, and validates your application's permissions.

Developers create a schema that models their permissions requirements and use a client library, such as this one, to apply the schema to the database, insert data into the database, and query the data efficiently to check permissions in their applications.

Supported client API versions:
- [v1alpha1](https://docs.authzed.com/reference/api#authzedapiv1alpha1)
- [v0](https://docs.authzed.com/reference/api#authzedapiv0)

You can find more info on each API on the [Authzed API reference documentation].
Additionally, Protobuf API documentation can be found on the [Buf Registry Authzed API repository].

[Authzed]: https://authzed.com
[Authzed API Reference documentation]: https://docs.authzed.com/reference/api
[Buf Registry Authzed API repository]: https://buf.build/authzed/api/docs/main

## Getting Started

We highly recommend following the **[Protecting Your First App]** guide to learn the latest best practice to integrate an application with Authzed.

If you're interested in examples for a specific version of the API, they can be found in their respective folders in the [examples directory].

[Protecting Your First App]: https://docs.authzed.com/guides/first-app
[examples directory]: /examples

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

[`NewClient()`]: https://pkg.go.dev/github.com/authzed/authzed-go/v0#NewClient
[Bearer Token]: https://datatracker.ietf.org/doc/html/rfc6750#section-2.1
[Authzed Dashboard]: https://app.authzed.com
[gRPC]: https://grpc.io
[DialOptions]: https://pkg.go.dev/google.golang.org/grpc?utm_source=godoc#DialOption

```go
import (
	"github.com/authzed/authzed-go/v0"
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
	"github.com/authzed/authzed-go/proto/authzed/api/v0"
	"github.com/authzed/authzed-go/v0"
	"github.com/authzed/grpcutil"
)

...

emilia := &v0.User{UserOneof: &v0.User_Userset{Userset: &v0.ObjectAndRelation{
	Namespace: "user",
	ObjectId:  "emilia",
	Relation:  "...",
}}}

post1Reader := &v0.ObjectAndRelation{Namespace: "post", ObjectId: "1", Relation: "read"}

// Is Emilia in the set of users that can read post #1?
resp, err := client.Check(ctx, &v0.CheckRequest{User: emilia, TestUserset: post1Reader})
if err != nil {
	log.Fatalf("failed to check permission: %s", err)
}

if resp.GetMembership() == v0.CheckResponse_MEMBER {
	log.Println("allowed!")
}
```
