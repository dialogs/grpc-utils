package util

import (
	"context"
	"time"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
)

// LogFromContext is a tiny wrapper to extract logger from context
func LogFromContext(ctx context.Context) *log.Entry {
	return grpc_logrus.Extract(ctx)
}

var durationLogger = grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
	return "grpc.time_ns", duration.Nanoseconds()
})
