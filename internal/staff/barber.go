package staff

import (
	"go-sleeping-barber/internal/customer"
	logger "go-sleeping-barber/pkg/log"
	"go.uber.org/zap"
	"time"
)

var log = logger.GetLogger()

type Barber struct {
	Sleeping    bool
	CutDuration time.Duration
}

func (b *Barber) CutHair(customer customer.Customer) {
	log.Info("Cutting hair for Customer", zap.String("customerID", customer.ID.String()))
	time.Sleep(b.CutDuration)
	log.Info("Haircut done for Customer", zap.String("customerID", customer.ID.String()))
}
