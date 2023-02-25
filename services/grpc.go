package services

import (
	"google.golang.org/grpc"
)

// Route gRPC into server adapter
func GrpcRouter(route *grpc.Server) {
	// protobuf.RegisterExampleServiceServer(route, &servers.GrpcExample{})
}
