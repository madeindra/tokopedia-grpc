package product

import (
	"bufio"
	"context"
	"io"
	"log"
	"strconv"
	"strings"

	appproto "github.com/Xanvial/tutorial-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	scanner bufio.Scanner
	client  appproto.ProductServiceClient // client implementation of grpc proto, update the name accordingly
}

func NewProductClient(
	reader io.Reader,
	// preferably add interceptor class as param here, but can also be added directly
) ProductClient {
	// initialize grpc and add interceptor here
	conn, err := grpc.Dial(":9000",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // to allow insecure connection
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return ProductClient{
		scanner: *bufio.NewScanner(reader),
		client:  appproto.NewProductServiceClient(conn),
	}
}

func (pc *ProductClient) MainLoop() {

	for {
		log.Println("---------------------")
		log.Println("Simple Client Command")
		log.Println("---------------------")
		log.Println("1) Add Product")
		log.Println("2) Get All Product")
		log.Println("3) Get Product by Id")
		log.Println("0) Exit")
		log.Print("->: ")
		// Scans a line from Stdin(Console)
		pc.scanner.Scan()
		// Holds the string that scanned
		text := strings.TrimSpace(pc.scanner.Text())

		switch text {
		case "1":
			pc.HandleAddProduct()
		case "2":
			pc.HandleGetProducts()
		case "3":
			pc.HandleGetProduct()
		case "0":
			return
		}
		log.Println("")
	}
}

func (pc *ProductClient) HandleAddProduct() {
	log.Println("Input product data")
	log.Println("---------------------")
	log.Println("id:")
	pc.scanner.Scan()
	productIDStr := strings.TrimSpace(pc.scanner.Text())
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		log.Println("id should be integer, err:", err)
		return
	}
	log.Println("name:")
	pc.scanner.Scan()
	productName := strings.TrimSpace(pc.scanner.Text())
	log.Println("description:")
	pc.scanner.Scan()
	productDesc := strings.TrimSpace(pc.scanner.Text())

	// Send this data to grpc server
	log.Println("add product with id:", productID, ", name:", productName, ", desc:", productDesc)

	resp, err := pc.client.AddProduct(context.Background(), &appproto.AddProductReq{
		Product: &appproto.Product{
			Id:          int64(productID),
			Name:        productName,
			Description: productDesc,
		},
	})

	if err != nil {
		log.Fatalf("Error when calling AddProduct: %s", err)
	}

	log.Printf("Response from server: %s", resp)
}

func (pc *ProductClient) HandleGetProducts() {
	// hit grpc server and print all response
	resp, err := pc.client.GetProducts(context.Background(), &appproto.GetProductsReq{})

	if err != nil {
		log.Fatalf("Error when calling GetProducts: %s", err)
	}

	log.Printf("Response from server: %s", resp)
}

func (pc *ProductClient) HandleGetProduct() {
	log.Println("Input product data")
	log.Println("---------------------")
	log.Println("id:")
	pc.scanner.Scan()
	productIDStr := strings.TrimSpace(pc.scanner.Text())
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		log.Println("id should be integer, err:", err)
		return
	}
	// Send this data to grpc server
	log.Println("request product id:", productID)

	// also print the response
	resp, err := pc.client.GetProduct(context.Background(), &appproto.GetProductReq{
		Id: int64(productID),
	})

	if err != nil {
		log.Fatalf("Error when calling GetProduct: %s", err)
	}

	log.Printf("Response from server: %s", resp)
}
