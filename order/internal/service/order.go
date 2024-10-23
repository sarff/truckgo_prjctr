package service

import (
	"context"

	shippingpb "github.com/alexandear/truckgo/shipping/grpc/grpcapi"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
)

const (
	TypeCustomer string = "customer" // when order was created
	TypeDriver   string = "driver"   // when order in status new so it can be accepted
)

func GetOrderPrice(ctx context.Context, client shippingpb.ShippingServiceClient, origin, destination string) (float64, error) {
	req := &shippingpb.PriceRequest{
		Origin:      origin,
		Destination: destination,
	}

	resp, err := client.CalculatePrice(ctx, req)
	if err != nil {
		return 0.0, err
	}

	return resp.GetPrice(), nil
}

func GetUserType(ctx context.Context, client userpb.UserServiceClient, userID uint32) (string, error) {
	req := &userpb.TypeRequest{
		UserId: userID,
	}

	resp, err := client.GetType(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetType(), nil
}
