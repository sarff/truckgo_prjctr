package main

import (
	"context"
	pb "github.com/alexandear/truckgo/order/grpcapi"
	"github.com/alexandear/truckgo/order/internal/models"
	"github.com/alexandear/truckgo/order/internal/repository"
	"github.com/alexandear/truckgo/order/internal/service"
	"github.com/alexandear/truckgo/shared/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedOrderServer

	log             *logging.Logger
	orderRepository *repository.Order
}

func (s *server) validateOrderID(orderID uint) (*models.Order, error) {
	order, err := s.orderRepository.FindOneByID(orderID)
	if err != nil {
		s.log.Error("Invalid order", "error", err)
		return nil, status.Error(codes.InvalidArgument, "invalid order id")
	}

	return &order, nil
}

// Create can be done only by user with type customer. The right user id and type are validated on frontend.
func (s *server) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	userID := uint(request.GetUserId())
	origin := request.GetOrigin()
	destination := request.GetDestination()

	price, err := service.GetOrderPrice(ctx, origin, destination)
	if err != nil {
		s.log.Error("Cannot calculate price", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	ID, err := s.orderRepository.Create(price, userID)
	if err != nil {
		s.log.Error("Cannot create order", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.log.Info("Created order", "id", ID)
	return &pb.CreateResponse{OrderId: uint32(ID)}, nil
}

// UpdateStatus can be done only by user with type driver. The right user id and type are validated on frontend.
func (s *server) UpdateStatus(_ context.Context, request *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	order, err := s.validateOrderID(uint(request.GetOrderId()))
	if err != nil {
		return nil, err
	}

	orderStatus := models.Status(request.GetStatus())
	// TODO validate status, StatusInProgress and StatusDone are valid here
	// TODO proper update
	err = s.orderRepository.UpdateStatus(*order, orderStatus)
	if err != nil {
		s.log.Error("Cannot update order status", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	//TODO figure out what is correct response ??
	return &pb.UpdateStatusResponse{}, nil
}

// Accept can be done only by user with type driver. The right user id and type are validated on frontend.
func (s *server) Accept(_ context.Context, request *pb.AcceptRequest) (*pb.AcceptResponse, error) {
	order, err := s.validateOrderID(uint(request.GetOrderId()))
	if err != nil {
		return nil, err
	}

	orderStatus := models.StatusAccepted
	// TODO validate status - should be after StatusNew and setDriverId to user_id
	// TODO proper update
	err = s.orderRepository.UpdateStatus(*order, orderStatus)
	if err != nil {
		s.log.Error("Cannot update order status", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	//TODO figure out what is correct response ??
	return &pb.AcceptResponse{}, nil
}

// Decline can be done by user with type driver or customer. The right user id and type are validated on frontend.
func (s *server) Decline(_ context.Context, request *pb.DeclineRequest) (*pb.DeclineResponse, error) {
	order, err := s.validateOrderID(uint(request.GetOrderId()))
	if err != nil {
		return nil, err
	}

	// TODO validate status - should be after StatusAccepted, sedDriverId to nil
	// TODO proper update
	orderStatus := models.StatusNew
	err = s.orderRepository.UpdateStatus(*order, orderStatus)
	if err != nil {
		s.log.Error("Cannot update order status", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	//TODO figure out what is correct response ??
	return &pb.DeclineResponse{}, nil
}

// Cancel can be done by user with type customer. The right user id and type are validated on frontend.
func (s *server) Cancel(_ context.Context, request *pb.CancelRequest) (*pb.CancelResponse, error) {
	order, err := s.validateOrderID(uint(request.GetOrderId()))
	if err != nil {
		return nil, err
	}

	// TODO validate status - should be after StatusNew
	// TODO proper update
	orderStatus := models.StatusCancelled
	err = s.orderRepository.UpdateStatus(*order, orderStatus)
	if err != nil {
		s.log.Error("Cannot update order status", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	//TODO figure out what is correct response ??
	return &pb.CancelResponse{}, nil
}

// Archive can be done by user with type customer. The right user id and type are validated on frontend.
func (s *server) Archive(_ context.Context, request *pb.ArchiveRequest) (*pb.ArchiveResponse, error) {
	//order, err := s.validateOrderID(uint(request.GetOrderId()))
	//if err != nil {
	//	return nil, err
	//}

	// TODO validate status - should be after StatusDone, isArchived set to true
	// TODO proper update
	//err = s.orderRepository.Update(*order)
	//if err != nil {
	//	s.log.Error("Cannot archive status", "error", err)
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	//TODO figure out what is correct response ??
	return &pb.ArchiveResponse{}, nil
}

// GetOne can be done by user with type customer or driver. The right user id and type are validated on frontend.
func (s *server) GetOne(_ context.Context, request *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	order, err := s.validateOrderID(uint(request.GetOrderId()))
	if err != nil {
		s.log.Error("Invalid request", "error", err)
		return nil, err
	}

	orderResponse := pb.OrderEntity{
		Number:   order.Number,
		Status:   pb.Status(order.Status),
		Price:    order.Price,
		UserId:   uint32(order.UserID),
		DriverId: uint32(order.DriverID),
	}

	return &pb.GetOneResponse{Order: &orderResponse}, nil
}

func (s *server) GetHistory(_ context.Context, request *pb.GetHistoryByUserRequest) (*pb.GetHistoryByUserResponse, error) {
	//TODO validation of user type
	return nil, nil
}

func (s *server) GetAll(_ context.Context, request *pb.GetAllByUserRequest) (*pb.GetAllByUserResponse, error) {
	//TODO validation of user type
	return nil, nil
}
