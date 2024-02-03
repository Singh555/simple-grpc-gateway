package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "grpcGateway/proto/helloworld"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := "localhost:8080" // gRPC server address
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
		return err
	}

	log.Println("HTTP server is running on :50051")
	return http.ListenAndServe(":50051", mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
