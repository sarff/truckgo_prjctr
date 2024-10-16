package main

import (
	"context"

	grpcapiUser "github.com/alexandear/truckgo/user-service/grpc/grpcapi"
)

type server struct {
	grpcapiUser.UnimplementedUserServiceServer
}

func (s *server) GetDrivers(context.Context, *grpcapiUser.DriverRequest) (*grpcapiUser.DriverResponse, error) {
	// TODO implement getting driver ids and positions

	// FIXME test currently just a test data
	stubDrivers := []*grpcapiUser.Driver {
		{
			Id: 1,
			Position: "метро Оболонь, Київ",
		},
		{
			Id: 2,
			Position: "метро Сирець, Київ",
		},
		{
			Id: 3,
			Position: "метро Нивки, Київ",
		},
		{
			Id: 4,
			Position: "Шпалерний ринок, Київ",
		},
		{
			Id: 5,
			Position: "метро Васильківська, Київ",
		},
		{
			Id: 6,
			Position: "метро Печерська, Київ",
		},
		{
			Id: 7,
			Position: "метро Позняки, Київ",
		},
		{
			Id: 8,
			Position: "метро Дарниця, Київ",
		},
		{
			Id: 9,
			Position: "Деснянський парк, Київ",
		},
	}

	return &grpcapiUser.DriverResponse{
		Drivers: stubDrivers,
	}, nil
}

func (s *server) TestFunc(context.Context, *grpcapiUser.TestRequest) (*grpcapiUser.TestResponse, error) {
	return &grpcapiUser.TestResponse{
		Message: "Some testing!",
	}, nil
}
