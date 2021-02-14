package main

import (
	"github.com/gogo/status"
	pb "github.com/psigen/bazel-go-issue/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	// Use gogo Status to test linking.
	foo := status.Error(codes.NotFound, "no such user")
	_ = foo

	// Create a client connection to the gRPC server.
	conn, _ := grpc.Dial("localhost:10000")
	client := pb.NewGreeterClient(conn)
	_ = client
}
