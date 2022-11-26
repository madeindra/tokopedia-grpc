package interceptor

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (gi *GRPCInterceptor) MetadataInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("cannot get metadata from context")
		}

		log.Printf("Metadata %v", md)

		return handler(ctx, req)
	}
}
