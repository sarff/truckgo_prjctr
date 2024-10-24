package services

import (
	"github.com/alexandear/truckgo/payment/internal/models"
	"github.com/alexandear/truckgo/shared/logging"
)

type PaymentGateway struct {
	pg *ThirdPartyPaymentGateway
	*logging.Logger
}

func NewPaymentGateway(logger *logging.Logger) *PaymentGateway {
	pg := &ThirdPartyPaymentGateway{Logger: logger}

	return &PaymentGateway{pg: pg, Logger: logger}
}

func (p *PaymentGateway) Pay(amount uint32, currency string, paymentMethod models.PaymentMethod) (uint32, string,
	error,
) {
	paymentID, status, err := p.pg.Payment(amount, currency, paymentMethod.Number)

	return paymentID, status, err
}

func (p *PaymentGateway) Payout(amount uint32, currency string, paymentMethod models.PaymentMethod) (uint32, string, error) {
	paymentID, status, err := p.pg.Payout(amount, currency, paymentMethod.Number)

	return paymentID, status, err
}
