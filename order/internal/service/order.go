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

const defaultDriversCount uint32 = 5

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

func GetUserPosition(ctx context.Context, client userpb.UserServiceClient, userID uint32) (string, error) {
	req := &userpb.TypeRequest{
		UserId: userID,
	}

	resp, err := client.GetType(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetType(), nil
}

func GetNearestDrivers(ctx context.Context, client shippingpb.ShippingServiceClient, userClient userpb.UserServiceClient, userID uint32, login string) ([]uint32, error) {
	user, err := userClient.GetUser(ctx, &userpb.UserRequest{Id: userID, Login: login})
	if err != nil {
		return nil, err
	}

	req := &shippingpb.DriverRequest{
		ClientLatitude:  user.Latitude,
		ClientLongitude: user.Longitude,
		DriversCount:    defaultDriversCount,
	}

	resp, err := client.FindTheNearestDrivers(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetDriverIds(), nil
}
