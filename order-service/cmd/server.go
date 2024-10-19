package main

import (
	"context"
	"github.com/alexandear/truckgo/order-service/grpcapi"
)

type server struct {
	grpcapi.UnimplementedOrderServiceServer
}

func (s *server) TestFunc(ctx context.Context, req *grpcapi.TestRequest) (*grpcapi.TestResponse, error) {
	return &grpcapi.TestResponse{
		Message: "Some testing!",
	}, nil
}
