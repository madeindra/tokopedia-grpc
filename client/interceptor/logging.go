package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func (gi *GRPCInterceptor) LoggingInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("Took %s", time.Since(start).String())

		return err
	}
}
