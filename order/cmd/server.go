package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/alexandear/truckgo/order/grpcapi"
	"github.com/alexandear/truckgo/order/internal/models"
	"github.com/alexandear/truckgo/order/internal/repository"
	"github.com/alexandear/truckgo/order/internal/service"
	"github.com/alexandear/truckgo/shared/logging"
	shippingpb "github.com/alexandear/truckgo/shipping/grpc/grpcapi"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
)

type server struct {
	pb.UnimplementedOrderServer

	log             *logging.Logger
	orderRepository *repository.Order
	shippingClient  shippingpb.ShippingServiceClient
	userClient      userpb.UserServiceClient
}

func (s *server) validateOrderID(orderID uint32) (*models.Order, error) {
	order, err := s.orderRepository.FindOneByID(orderID)
	if err != nil {
		s.log.Error("Invalid order", "error", err)
		return nil, status.Error(codes.InvalidArgument, "invalid order id")
	}

	return &order, nil
}

// Create can be done only by user with type customer. The right user id and type are validated on frontend.
func (s *server) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	userID := request.GetUserId()
	origin := request.GetOrigin()
	destination := request.GetDestination()

	price, err := service.GetOrderPrice(ctx, s.shippingClient, origin, destination)
	if err != nil {
		s.log.Error("Cannot calculate price", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	id, err := s.orderRepository.Create(price, userID)
	if err != nil {
		s.log.Error("Cannot create order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.log.Info("Created order", "id", id)
	return &pb.CreateResponse{OrderId: id}, nil
}

// UpdateStatus can be done only by user with type driver. The right user id and type are validated on frontend.
func (s *server) UpdateStatus(_ context.Context, request *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	order, err := s.validateOrderID(request.GetOrderId())
	if err != nil {
		return nil, err
	}

	newStatus := models.Status(request.GetStatus())
	isValid := models.ValidateStatus(*order, newStatus)
	if !isValid {
		err = status.Errorf(codes.FailedPrecondition, "Cannot change order status to %d", newStatus)
		s.log.Error("Invalid order status", "error", err)
		return nil, err
	}

	updates := map[string]any{
		"Status": newStatus,
	}
	err = s.orderRepository.Update(*order, updates)
	if err != nil {
		s.log.Error("Cannot update order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateStatusResponse{}, nil
}

// Accept can be done only by user with type driver. The right user id and type are validated on frontend.
func (s *server) Accept(_ context.Context, request *pb.AcceptRequest) (*pb.AcceptResponse, error) {
	order, err := s.validateOrderID(request.GetOrderId())
	if err != nil {
		return nil, err
	}

	newStatus := models.StatusAccepted
	isValid := models.ValidateStatus(*order, newStatus)
	if !isValid {
		err = status.Errorf(codes.FailedPrecondition, "Cannot change order status to %d", newStatus)
		s.log.Error("Invalid order status", "error", err)
		return nil, err
	}

	updates := map[string]any{
		"Status":   newStatus,
		"DriverID": request.GetUserId(),
	}
	err = s.orderRepository.Update(*order, updates)
	if err != nil {
		s.log.Error("Cannot update order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AcceptResponse{}, nil
}

// Decline can be done by user with type driver or customer. The right user id and type are validated on frontend.
func (s *server) Decline(_ context.Context, request *pb.DeclineRequest) (*pb.DeclineResponse, error) {
	order, err := s.validateOrderID(request.GetOrderId())
	if err != nil {
		return nil, err
	}

	newStatus := models.StatusNew
	isValid := models.ValidateStatus(*order, newStatus)
	if !isValid {
		err = status.Errorf(codes.FailedPrecondition, "Cannot change order status to %d", newStatus)
		s.log.Error("Invalid order status", "error", err)
		return nil, err
	}

	updates := map[string]any{
		"Status":   newStatus,
		"DriverID": nil,
	}
	err = s.orderRepository.Update(*order, updates)
	if err != nil {
		s.log.Error("Cannot update order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeclineResponse{}, nil
}

// Cancel can be done by user with type customer. The right user id and type are validated on frontend.
func (s *server) Cancel(_ context.Context, request *pb.CancelRequest) (*pb.CancelResponse, error) {
	order, err := s.validateOrderID(request.GetOrderId())
	if err != nil {
		return nil, err
	}

	newStatus := models.StatusCancelled
	isValid := models.ValidateStatus(*order, newStatus)
	if !isValid {
		err = status.Errorf(codes.FailedPrecondition, "Cannot change order status to %d", newStatus)
		s.log.Error("Invalid order status", "error", err)
		return nil, err
	}

	updates := map[string]any{
		"Status": newStatus,
	}
	err = s.orderRepository.Update(*order, updates)
	if err != nil {
		s.log.Error("Cannot update order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CancelResponse{}, nil
}

// Archive can be done by user with type customer. The right user id and type are validated on frontend.
func (s *server) Archive(_ context.Context, request *pb.ArchiveRequest) (*pb.ArchiveResponse, error) {
	order, err := s.validateOrderID(request.GetOrderId())
	if err != nil {
		return nil, err
	}

	if order.Status != models.StatusDone {
		err = status.Error(codes.FailedPrecondition, "In order to archive order should be done")
		s.log.Error("Invalid order status", "error", err)
		return nil, err
	}

	updates := map[string]any{
		"IsArchived": true,
	}
	err = s.orderRepository.Update(*order, updates)
	if err != nil {
		s.log.Error("Cannot update order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ArchiveResponse{}, nil
}

// GetOne can be done by user with type customer or driver. The right user id and type are validated on frontend.
func (s *server) GetOne(_ context.Context, request *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	order, err := s.validateOrderID(request.GetOrderId())
	if err != nil {
		s.log.Error("Invalid request", "error", err)
		return nil, err
	}

	orderResponse := pb.OrderEntity{
		Number:     order.Number,
		Status:     pb.Status(order.Status),
		Price:      order.Price,
		UserId:     order.UserID,
		DriverId:   order.DriverID,
		IsArchived: order.IsArchived,
	}

	return &pb.GetOneResponse{Order: &orderResponse}, nil
}

func (s *server) GetHistoryByUser(ctx context.Context, request *pb.GetHistoryByUserRequest) (*pb.GetHistoryByUserResponse, error) {
	page := int(request.GetPage())
	limit := int(request.GetLimit())
	userID := request.GetUserId()
	userType, err := service.GetUserType(ctx, s.userClient, userID)
	if err != nil {
		s.log.Error("Cannot validate user type", "error", err)
		return nil, err
	}

	filters := make(map[string]any)

	switch userType {
	case service.TypeCustomer:
		filters["user_id"] = userID
	case service.TypeDriver:
		filters["driver_id"] = userID
	}

	filters["is_archived"] = request.GetIsArchived()
	filters["status"] = models.StatusDone

	orders, total, err := s.orderRepository.FindAll(page, limit, filters)
	if err != nil {
		s.log.Error("Cannot get orders from database", "error", err)
		return nil, err
	}

	ordersResponse := make([]*pb.OrderEntity, 0, len(orders))
	for _, order := range orders {
		ordersResponse = append(ordersResponse, &pb.OrderEntity{
			Number:     order.Number,
			Status:     pb.Status(order.Status),
			Price:      order.Price,
			UserId:     order.UserID,
			DriverId:   order.DriverID,
			IsArchived: order.IsArchived,
		})
	}

	return &pb.GetHistoryByUserResponse{Orders: ordersResponse, Total: total}, nil
}

func (s *server) GetAllByUser(ctx context.Context, request *pb.GetAllByUserRequest) (*pb.GetAllByUserResponse, error) {
	page := int(request.GetPage())
	limit := int(request.GetLimit())
	userID := request.GetUserId()
	userType, err := service.GetUserType(ctx, s.userClient, userID)
	if err != nil {
		s.log.Error("Cannot validate user type", "error", err)
		return nil, err
	}

	filters := make(map[string]any)

	switch userType {
	case service.TypeCustomer:
		filters["user_id"] = userID
	case service.TypeDriver:
		filters["driver_id"] = userID
	}

	if statusWrapper, ok := request.GetOptionalStatus().(*pb.GetAllByUserRequest_Status); ok {
		filters["status"] = models.Status(statusWrapper.Status)
	}

	orders, total, err := s.orderRepository.FindAll(page, limit, filters)
	if err != nil {
		s.log.Error("Cannot get orders from database", "error", err)
		return nil, err
	}

	ordersResponse := make([]*pb.OrderEntity, 0, len(orders))
	for _, order := range orders {
		ordersResponse = append(ordersResponse, &pb.OrderEntity{
			Number:     order.Number,
			Status:     pb.Status(order.Status),
			Price:      order.Price,
			UserId:     order.UserID,
			DriverId:   order.DriverID,
			IsArchived: order.IsArchived,
		})
	}

	return &pb.GetAllByUserResponse{Orders: ordersResponse, Total: total}, nil
}
