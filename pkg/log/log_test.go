package log

import (
	"context"
	"golang.org/x/exp/slog"
	"testing"
)

func Test_t(t *testing.T) {
	ctx := context.Background()
	slog.Info("Info to Ethereum")
	ctx = context.WithValue(ctx, "key", "value")
	slog.InfoContext(ctx, "Info to Ethereum")
}
