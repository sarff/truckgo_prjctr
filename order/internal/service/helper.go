package service

import (
	"context"
	shippingpb "github.com/alexandear/truckgo/shipping-service/grpc/grpcapi"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	TypeCustomer string = "customer" // when order was created
	TypeDriver   string = "driver"   // when order in status new so it can be accepted
)

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

func GetUserType(ctx context.Context, userID uint) (string, error) {
	conn, err := grpc.NewClient("user-server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	userClient := userpb.NewUserServiceClient(conn)

	req := &userpb.TypeRequest{
		UserId: uint32(userID),
	}

	resp, err := userClient.GetType(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetType(), nil
}
