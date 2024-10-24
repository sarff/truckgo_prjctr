package main

import (
	"context"
	"errors"

	pb "github.com/alexandear/truckgo/payment/grpcapi"
	"github.com/alexandear/truckgo/payment/internal/repository"
	"github.com/alexandear/truckgo/payment/internal/services"
	"github.com/alexandear/truckgo/shared/logging"
)

type server struct {
	pb.UnimplementedPaymentServer
	pg                      *services.PaymentGateway
	paymentRepository       *repository.Payment
	paymentMethodRepository *repository.PaymentMethod
	log                     *logging.Logger
}

func (s *server) Pay(ctx context.Context, request *pb.PayRequest) (*pb.PayResponse, error) {
	userID := request.GetUserId()
	driverID := request.GetDriverId()
	amount := request.GetAmount()
	currency := request.GetCurrency()
	s.log.Debug("Pay request data",
		"userID", userID,
		"driverID", driverID,
		"amount", amount,
		"currency", currency,
	)

	driverPaymentMethod, err := s.paymentMethodRepository.FindOneByUser(driverID)
	if err != nil {
		return nil, errors.New("driver payment method not found")
	}

	userPaymentMethod, err := s.paymentMethodRepository.FindOneByUser(userID)
	if err != nil {
		return nil, errors.New("user payment method not found")
	}

	externalID, paymentStatus, err := s.pg.Pay(amount, currency, userPaymentMethod)
	if err != nil {
		return nil, err
	}

	if paymentStatus != "success" && externalID == 0 {
		return nil, errors.New("payment failed")
	}

	var paymentID uint32
	if paymentStatus == "success" {
		paymentID, err = s.paymentRepository.CreateSuccessful(amount, currency, externalID, userPaymentMethod.ID, userID)
	} else {
		paymentID, err = s.paymentRepository.CreateFail(amount, currency, externalID, userPaymentMethod.ID, userID)
	}

	if err != nil {
		s.log.Error("Payment create error", "error", err.Error(), "externalID", externalID)
	}
	// payout money to driver and minus a 10% service fee
	payoutID, payoutStatus, err := s.pg.Payout((amount*90)/100, currency, driverPaymentMethod)
	if err != nil {
		s.log.Info("Payout failed: %s", err.Error())
	}
	if payoutStatus != "success" {
		s.log.Info("Payout failed: paymentID: %d", paymentID)
	}

	s.log.Info("Payout successfully created: payoutID: %d", payoutID)
	s.log.Info("Payment successfully created: paymentID: %d", paymentID)

	return &pb.PayResponse{PaymentId: paymentID}, nil
}

func (s *server) CreatePaymentMethod(ctx context.Context, request *pb.CreatePaymentMethodRequest) (*pb.
	CreatePaymentMethodResponse, error,
) {
	userID := request.GetUserId()
	number := request.GetCardNumber()
	cvv := request.GetCardCvv()
	expiry := request.GetCardExpiration()
	name := request.GetCardHolderName()

	paymentMethodID, err := s.paymentMethodRepository.Create(userID, number, cvv, expiry, name)

	if paymentMethodID == 0 || err != nil {
		s.log.Error("Payment method not created", "error", err.Error())

		return nil, errors.New("payment method not created")
	}

	return &pb.CreatePaymentMethodResponse{PaymentMethodId: paymentMethodID}, nil
}

func (s *server) DeletePaymentMethod(ctx context.Context, request *pb.DeletePaymentMethodRequest) (*pb.DeletePaymentMethodResponse,
	error,
) {
	paymentMethodID := request.GetPaymentMethodId()
	s.log.Debug("Delete payment method request", "paymentMethodID", paymentMethodID)

	paymentMethod, err := s.paymentMethodRepository.FindByID(paymentMethodID)
	if err != nil {
		s.log.Error("Payment method not found", "paymentMethodID", paymentMethodID)

		return nil, errors.New("payment method not found")
	}

	err = s.paymentMethodRepository.Deactivate(paymentMethod.ID)
	if err != nil {
		s.log.Error("Payment method not deleted", "error", err.Error())

		return nil, errors.New("payment method not deleted")
	}

	return &pb.DeletePaymentMethodResponse{Message: "Successfully delete"}, nil
}
