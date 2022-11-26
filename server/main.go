package main

import (
	"log"
	"net"

	appproto "github.com/Xanvial/tutorial-grpc/proto"
	"github.com/Xanvial/tutorial-grpc/server/handler"
	"github.com/Xanvial/tutorial-grpc/server/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// use this product class Usecase as param for grpc handler
	productData := usecase.NewProductUsecase()

	// put grpc class initialization and interceptor here
	productHandler := handler.NewProductHandler(productData)

	grpcServer := grpc.NewServer()

	appproto.RegisterProductServiceServer(grpcServer, productHandler)

	// register server using reflection
	reflection.Register(grpcServer)

	log.Println("start listening on port: 9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
