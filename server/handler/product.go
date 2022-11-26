package handler

import (
	"context"

	appproto "github.com/Xanvial/tutorial-grpc/proto"
	"github.com/Xanvial/tutorial-grpc/server/model"
	"github.com/Xanvial/tutorial-grpc/server/usecase"
)

type ProductServer struct {
	productUC usecase.ProductClass

	// UnsafeProductServiceServer is used for grpc forward-compatibility.
	//
	// in the case of adding new method in proto file, there's two approach for forward-compatibility:
	//
	// 1. like in example/server/hello/hello.go, adding UnimplementedXXXX will always make the compile successfull
	// even if there's no implementation for new method.
	// But on runtime, if the new method is called, it will automatically return error.
	//
	// 2. this approach, adding UnsafeXXXX will force compile error if new method is not implemented

	// uncomment this, after correctly import the proto file
	appproto.UnsafeProductServiceServer
}

func NewProductHandler(productUsecase usecase.ProductClass) *ProductServer {
	return &ProductServer{
		productUC: productUsecase,
	}
}

// Put all other grpc handlers in here
func (s *ProductServer) AddProduct(ctx context.Context, in *appproto.AddProductReq) (*appproto.AddProductResp, error) {
	err := s.productUC.AddProduct(model.Product{
		ID:          int(in.Product.GetId()),
		Name:        in.Product.GetName(),
		Description: in.Product.GetDescription(),
	})

	return &appproto.AddProductResp{
		Success: true,
	}, err
}

func (s *ProductServer) GetProducts(ctx context.Context, in *appproto.GetProductsReq) (*appproto.GetProductsResp, error) {
	products := s.productUC.GetProducts()

	resp := &appproto.GetProductsResp{}

	for _, val := range products {
		resp.Products = append(resp.Products, &appproto.Product{
			Id:          int64(val.ID),
			Name:        val.Name,
			Description: val.Description,
		})
	}

	return resp, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, in *appproto.GetProductReq) (*appproto.GetProductResp, error) {
	data, err := s.productUC.GetProduct(int(in.GetId()))
	if err != nil {
		return nil, err
	}

	resp := &appproto.GetProductResp{
		Product: &appproto.Product{
			Id:          int64(data.ID),
			Name:        data.Name,
			Description: data.Description,
		},
	}

	return resp, nil
}
