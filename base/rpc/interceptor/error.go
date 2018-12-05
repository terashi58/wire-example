package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// ErrorUnaryServerInterceptor returns a new unary server interceptor that handle errors.
func ErrorUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		if err != nil {
			log.Printf("%+v", err)
		}

		return res, err
	}
}
