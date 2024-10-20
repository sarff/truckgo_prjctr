package services

import (
	"context"
	"fmt"
	"github.com/alexandear/truckgo/shared/logging"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
	"github.com/alexandear/truckgo/user/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"regexp"
	"time"
)

type UserServiceServer struct {
	*gorm.DB
	userpb.UnimplementedUserServiceServer
	*logging.Logger
}

func (u *UserServiceServer) GetDrivers(context.Context, *userpb.DriverRequest) (*userpb.DriverResponse, error) {
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

var regexPhone = regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`)

func (u *UserServiceServer) NewCustomer(_ context.Context, req *userpb.NewCustomerRequest) (*userpb.NewCustomerResponse, error) {
	matches := regexPhone.FindAllString(req.Phone, -1)
	if len(matches) == 0 {
		u.Logger.Error("invalid phone format", logging.ErrInvalidPhone, req.Phone)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid phone format %s", req.Phone))
	}

	newUser := &models.User{
		Login:      req.Login,
		FullName:   req.FullName,
		TypeUserID: 1,
		Phone:      req.Phone,
		CreatedAt:  time.Now(),
	}

	if err := u.DB.Create(&newUser).Error; err != nil {
		u.Logger.Error("failed to create user", logging.ErrDBCreateFailed, err)
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &userpb.NewCustomerResponse{
		Message: fmt.Sprintf("User created"),
	}, nil
}

func (u *UserServiceServer) GetCustomer(_ context.Context, req *userpb.GetCustomerRequest) (*userpb.GetCustomerResponse, error) {

	return &userpb.GetCustomerResponse{
		Customers: nil,
	}, nil
}

func (u *UserServiceServer) NewDriver(_ context.Context, req *userpb.NewDriverRequest) (*userpb.NewDriverResponse, error) {
	matches := regexPhone.FindAllString(req.Phone, -1)
	if len(matches) == 0 {
		u.Logger.Error("invalid phone format", logging.ErrInvalidPhone, req.Phone)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid phone format %s", req.Phone))
	}

	newUser := &models.User{
		Login:      req.Login,
		FullName:   req.FullName,
		TypeUserID: 2,
		Phone:      req.Phone,
		CreatedAt:  time.Now(),
	}

	newDriver := &models.Driver{
		User:      *newUser,
		License:   req.License,
		CarModel:  req.CarModel,
		CarNumber: req.CarNumber,
	}

	if err := u.DB.Create(&newDriver).Error; err != nil {
		u.Logger.Error("failed to create user", logging.ErrDBCreateFailed, err)
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &userpb.NewDriverResponse{
		Message: fmt.Sprintf("User created"),
	}, nil
}

func (u *UserServiceServer) UpdateUser(_ context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return &userpb.UpdateUserResponse{
		Message: fmt.Sprintf("User with phone %s - updated", req.Login),
	}, nil
}
func (u *UserServiceServer) GetType(_ context.Context, req *userpb.TypeRequest) (*userpb.TypeResponse, error) {
	var user *models.User

	var typeUser *models.TypeUser

	if err := u.DB.First(&user, req.UserId).Error; err != nil {
		u.Logger.Info("Record not Found", "GetType_User", err)
		return nil, status.Error(codes.NotFound, "Record not found")
	}

	if err := u.DB.First(&typeUser, user.TypeUserID).Error; err != nil {
		u.Logger.Info("Record not Found", "GetType_Type", err)
		return nil, status.Error(codes.NotFound, "Record not found")
	}

	return &userpb.TypeResponse{
		Type:    typeUser.Type,
		Message: fmt.Sprintf("User with Login %v has type %s", req.UserId, typeUser.Type),
	}, nil
}

func (u *UserServiceServer) GetUser(_ context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	return &userpb.UserResponse{
		Message: fmt.Sprintf("success"),
	}, nil
}
