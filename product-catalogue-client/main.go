// https://grpc.io/docs/languages/go/quickstart/
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	// pb "pb."
	pb "kszpakowski.com/grpc-client/pb"
)

const (
	defaultId = 1
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	id   = flag.Int("id", defaultId, "Product id to fetch details")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductCatalogClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetProductDetails(ctx, &pb.ProductDetailsRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Product Details: %s", r.GetName())
}
