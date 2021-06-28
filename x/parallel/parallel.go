// Package parallel implements experimental utilities for performing parallel
// client interactions with the Authzed API.
package parallel

import (
	"context"

	"google.golang.org/grpc"

	"github.com/authzed/authzed-go/internal/ctxgroup"
	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
)

// Check performs the provided list of CheckRequests in parallel and returns
// a slice of responses that are stored at their respective indices.
func Check(ctx context.Context, client v0.ACLServiceClient, in []*v0.CheckRequest, opts ...grpc.CallOption) ([]*v0.CheckResponse, error) {
	results := make([]*v0.CheckResponse, len(in))
	g := ctxgroup.WithContext(ctx)
	for i, request := range in {
		i, request := i, request // close over the value that will not change
		g.GoCtx(func(ctx context.Context) (err error) {
			results[i], err = client.Check(ctx, request, opts...)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

// ContentChangeCheck performs the provided list of ContentChangeCheckRequests
// in parallel and returns a slice of responses that are stored at their
// respective indices.
func ContentChangeCheck(ctx context.Context, client v0.ACLServiceClient, in []*v0.ContentChangeCheckRequest, opts ...grpc.CallOption) ([]*v0.CheckResponse, error) {
	results := make([]*v0.CheckResponse, len(in))
	g := ctxgroup.WithContext(ctx)
	for i, request := range in {
		i, request := i, request // close over the value that will not change
		g.GoCtx(func(ctx context.Context) (err error) {
			results[i], err = client.ContentChangeCheck(ctx, request, opts...)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

// AllChecksAreMember returns true if every provided CheckResponse has
// membership.
func AllChecksAreMember(resps []*v0.CheckResponse) bool {
	for _, resp := range resps {
		if resp.Membership != v0.CheckResponse_MEMBER {
			return false
		}
	}
	return true
}
