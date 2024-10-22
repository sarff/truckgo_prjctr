package services

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"

	"github.com/alexandear/truckgo/shared/logging"
)

type ThirdPartyPaymentGateway struct {
	*logging.Logger
}

func (p *ThirdPartyPaymentGateway) Payment(amount uint32, currency string, card string) (id uint32, status string,
	err error,
) {
	isSuccess := randUint32(0, 10) != 0
	if isSuccess {
		id = randUint32(0, math.MaxUint32)
		p.Logger.Info("Make fake successful Payment", "id", id, "card", card, "amount", amount, "currency", currency)

		return id, "success", nil
	}

	err = errors.New("incufficient funds")

	return 0, "fail", err
}

func (p *ThirdPartyPaymentGateway) Payout(amount uint32, currency string, card string) (id uint32, status string, err error) {
	isSuccess := randUint32(0, 10) != 0
	if isSuccess {
		id = randUint32(0, math.MaxUint32)
		p.Logger.Info("Make fake successful Payout", "id", id, "card", card, "amount", amount, "currency", currency)

		return id, "success", nil
	}

	err = errors.New("expired card")

	return 0, "fail", err
}

func (p *ThirdPartyPaymentGateway) Refund(id int) error {
	isSuccess := randUint32(0, 10) != 0
	if isSuccess {
		p.Logger.Info("Make fake successful Refund %d", id)

		return nil
	}

	return errors.New("refund failed")
}

func randUint32(min, max uint32) uint32 {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	number := n.Int64()
	if number < 0 {
		return 0
	}
	if number > math.MaxUint32 {
		return 0
	}

	return min + uint32(number)
}
