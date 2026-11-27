package requestmeta

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestWithDispatchStrategy(t *testing.T) {
	t.Parallel()

	expr := "prefer_secondary(['tiger'], duration('50ms'))"
	ctx := WithDispatchStrategy(context.Background(), expr)

	md, ok := metadata.FromOutgoingContext(ctx)
	require.True(t, ok)
	require.Equal(t, []string{expr}, md.Get(string(RequestDispatchStrategyKey)))
}
