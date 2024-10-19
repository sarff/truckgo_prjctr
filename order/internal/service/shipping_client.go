package service

import (
	"context"
	shippingpb "github.com/alexandear/truckgo/shipping-service/grpc/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO move connection to server
func GetOrderPrice(ctx context.Context, origin, destination string) (float64, error) {
	conn, err := grpc.NewClient("auth-server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0.0, err
	}
	defer conn.Close()

	shippingClient := shippingpb.NewShippingServiceClient(conn)

	req := &shippingpb.PriceRequest{
		Origin:      origin,
		Destination: destination,
	}

	resp, err := shippingClient.CalculatePrice(ctx, req)
	if err != nil {
		return 0.0, err
	}

	return resp.GetPrice(), nil
}
