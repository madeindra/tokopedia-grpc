package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func (gi *GRPCInterceptor) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("LoggingInterceptor: %T \n", req)

		return handler(ctx, req)
	}
}
