package main

import (
	"context"

	grpcapiOrder "github.com/alexandear/truckgo/order-service/grpc/grpcapi"
)

type server struct {
	grpcapiOrder.UnimplementedOrderServiceServer
}

func (s *server) TestFunc(ctx context.Context, req *grpcapiOrder.TestRequest) (*grpcapiOrder.TestResponse, error) {
	return &grpcapiOrder.TestResponse{
		Message: "Some testing!",
	}, nil
}
