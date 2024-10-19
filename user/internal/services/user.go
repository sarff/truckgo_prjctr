package services

import (
	"context"
	"github.com/alexandear/truckgo/shared/logging"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	*gorm.DB
	userpb.UnimplementedUserServiceServer
	*logging.Logger
}

func (s *UserServiceServer) GetDrivers(context.Context, *userpb.DriverRequest) (*userpb.DriverResponse, error) {
	// TODO implement getting driver ids and positions

	// FIXME test currently just a test data
	stubDrivers := []*userpb.Driver{
		{
			Id:       1,
			Position: "метро Оболонь, Київ",
		},
		{
			Id:       2,
			Position: "метро Сирець, Київ",
		},
		{
			Id:       3,
			Position: "метро Нивки, Київ",
		},
		{
			Id:       4,
			Position: "Шпалерний ринок, Київ",
		},
		{
			Id:       5,
			Position: "метро Васильківська, Київ",
		},
		{
			Id:       6,
			Position: "метро Печерська, Київ",
		},
		{
			Id:       7,
			Position: "метро Позняки, Київ",
		},
		{
			Id:       8,
			Position: "метро Дарниця, Київ",
		},
		{
			Id:       9,
			Position: "Деснянський парк, Київ",
		},
	}

	return &userpb.DriverResponse{
		Drivers: stubDrivers,
	}, nil
}
