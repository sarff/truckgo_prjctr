package services

import (
	"context"
	"github.com/alexandear/truckgo/shared/logging"
	pb "github.com/alexandear/truckgo/user/grpcapi"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	*gorm.DB
	pb.UnimplementedUserServiceServer
	*logging.Logger
}

func (s *UserServiceServer) GetDrivers(context.Context, *pb.DriverRequest) (*pb.DriverResponse, error) {
	// TODO implement getting driver ids and positions

	// FIXME test currently just a test data
	stubDrivers := []*pb.Driver{
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

	return &pb.DriverResponse{
		Drivers: stubDrivers,
	}, nil
}

func (s *UserServiceServer) TestFunc(context.Context, *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Message: "Some testing!",
	}, nil
}
