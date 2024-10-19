package main

import (
	"context"
	pb "github.com/alexandear/truckgo/order/grpcapi"
	"github.com/alexandear/truckgo/order/internal/repository"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedOrderServer

	db              *gorm.DB
	orderRepository *repository.Order
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *server) UpdateStatus(ctx context.Context, in *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	return nil, nil
}

func (s *server) Accept(ctx context.Context, in *pb.AcceptRequest) (*pb.AcceptResponse, error) {
	return nil, nil
}

func (s *server) Decline(ctx context.Context, in *pb.DeclineRequest) (*pb.DeclineResponse, error) {
	return nil, nil
}

func (s *server) Cancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelResponse, error) {
	return nil, nil
}

func (s *server) Archive(ctx context.Context, in *pb.ArchiveRequest) (*pb.ArchiveResponse, error) {
	return nil, nil
}

func (s *server) GetOne(ctx context.Context, in *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	return nil, nil
}

func (s *server) GetHistory(ctx context.Context, in *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error) {
	return nil, nil
}

func (s *server) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	return nil, nil
}
