package interceptor

import (
	"google.golang.org/grpc"
)

// Default returns a set of default gRPC interceptors for a unary server.
func Default() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{}
}
