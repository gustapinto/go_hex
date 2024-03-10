package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func Log(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
	start := time.Now()

	res, err = handler(ctx, req)

	slog.Info("", "method", info.FullMethod, "duration", time.Since(start))
	return
}
