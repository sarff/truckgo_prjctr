package services

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/alexandear/truckgo/shared/logging"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
	"github.com/alexandear/truckgo/user/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	*gorm.DB
	userpb.UnimplementedUserServiceServer
	*logging.Logger
}

func convertUserToDriver(user *models.User) *userpb.Driver {
	return &userpb.Driver{
		Id:        user.ID,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}
}

func convertUserToCustomer(user *models.User) *userpb.Customer {
	return &userpb.Customer{
		Id:        user.ID,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}
}

func (u *UserServiceServer) ListDrivers(_ context.Context, req *userpb.ListDriverRequest) (*userpb.ListDriverResponse, error) {
	var drivers []*userpb.Driver
	var users []*models.User
	err := u.DB.Select("id", "latitude", "longitude").
		Where("status = ? AND type_user_id = ?", true, 2).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		driver := convertUserToDriver(user)
		drivers = append(drivers, driver)
	}

	return &userpb.ListDriverResponse{
		Drivers: drivers,
	}, nil
}

var regexPhone = regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`)

func (u *UserServiceServer) NewCustomer(_ context.Context, req *userpb.NewCustomerRequest) (*userpb.NewCustomerResponse, error) {
	if req.Phone != "" {
		matches := regexPhone.FindAllString(req.Phone, -1)
		if len(matches) == 0 {
			u.Logger.Error("invalid phone format", logging.ErrInvalidPhone, req.Phone)
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid %s phone format", req.Phone))
		}
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
		Message: "user created",
	}, nil
}

func (u *UserServiceServer) ListCustomers(_ context.Context, req *userpb.ListCustomerRequest) (*userpb.ListCustomerResponse, error) {
	var customers []*userpb.Customer
	var users []*models.User
	err := u.DB.Select("id", "latitude", "longitude").
		Where("status = ? AND type_user_id = ?", true, 1).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		customer := convertUserToCustomer(user)
		customers = append(customers, customer)
	}

	return &userpb.ListCustomerResponse{
		Customers: customers,
	}, nil
}

func (u *UserServiceServer) NewDriver(_ context.Context, req *userpb.NewDriverRequest) (*userpb.NewDriverResponse, error) {
	if req.Phone != "" {
		matches := regexPhone.FindAllString(req.Phone, -1)
		if len(matches) == 0 {
			u.Logger.Error("invalid phone format", logging.ErrInvalidPhone, req.Phone)
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid %s phone format", req.Phone))
		}
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
		Message: "User created",
	}, nil
}

func (u *UserServiceServer) UpdateUser(_ context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	var updateUser *models.User

	if err := u.DB.First(&updateUser, req.Id).Error; err != nil {
		u.Logger.Error("failed to update user", logging.ErrDBUpdateFailed, err)
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	updateUser.FullName = req.FullName
	updateUser.Status = req.Status
	updateUser.Phone = req.Phone
	updateUser.Rating = req.Rating
	updateUser.Latitude = req.Latitude
	updateUser.Longitude = req.Longitude

	if err := u.DB.Save(&updateUser).Error; err != nil {
		u.Logger.Error("failed to update user", logging.ErrDBUpdateFailed, err)
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return &userpb.UpdateUserResponse{
		Message: fmt.Sprintf("User with phone %v - updated", req.Id),
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
	var user *models.User

	if err := u.DB.First(&user, req.Id).Error; err != nil {
		if err := u.DB.First(&user, req.Login).Error; err != nil {
			u.Logger.Info("Record not Found", "GetUser", err)
			return nil, status.Error(codes.NotFound, "Record not found")
		}
	}

	return &userpb.UserResponse{
		Id:         user.ID,
		Login:      user.Login,
		FullName:   user.FullName,
		TypeUserID: user.TypeUserID,
		Status:     user.Status,
		Phone:      user.Phone,
		Rating:     user.Rating,
		Latitude:   user.Latitude,
		Longitude:  user.Longitude,
	}, nil
}
