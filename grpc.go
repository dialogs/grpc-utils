package util

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// DefaultGrpcOptions provides basic gRPC-server infrastructure
func DefaultGrpcOptions(
	logger *log.Entry,
	unaryInterceptors []grpc.UnaryServerInterceptor,
	streamInterceptors []grpc.StreamServerInterceptor) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logger, durationLogger),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_middleware.ChainUnaryServer(unaryInterceptors...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logger, durationLogger),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_middleware.ChainStreamServer(streamInterceptors...),
		)}
}
