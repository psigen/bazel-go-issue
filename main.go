package main

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func main() {
	foo := status.Error(codes.NotFound, "no such user")
	_ = foo
	return
}
