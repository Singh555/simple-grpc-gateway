## gRPC Gateway with Go

This Go program demonstrates the setup of a gRPC Gateway, which serves as an HTTP/REST proxy for a gRPC server. The code utilizes the grpc-gateway library to generate a reverse proxy that translates RESTful HTTP requests into gRPC calls.
How it Works

    The run function sets up a new HTTP server and a grpc-gateway runtime ServeMux.
    It establishes a connection to the gRPC server running on localhost:8080 using the pb.RegisterGreeterHandlerFromEndpoint function.
    HTTP requests are handled by the runtime ServeMux, which forwards them to the gRPC server.

# How to Run

To run the code, follow these steps:

```bash

# Clone the repository
git clone https://github.com/your-username/your-repository.git
cd your-repository

# Install dependencies
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/runtime
go get -u google.golang.org/grpc

# Run the code
go run main.go
```
Ensure that you have Go installed on your machine.
Access the gRPC Gateway

The gRPC Gateway is accessible at localhost:50051. It translates incoming HTTP requests into gRPC calls, forwarding them to the gRPC server running on localhost:8080.
Dependencies

The code relies on the following external dependencies:

    github.com/grpc-ecosystem/grpc-gateway/v2/runtime: A library for building gRPC-to-HTTP gateways.
    google.golang.org/grpc: The official gRPC Go implementation.