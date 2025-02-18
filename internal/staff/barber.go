package staff

import (
	"github.com/google/uuid"
	"go-sleeping-barber/internal/sleep"
	logger "go-sleeping-barber/pkg/log"
	"go.uber.org/zap"
	"time"
)

var log = logger.GetLogger()

type Barber struct {
	Sleeping    bool
	CutDuration time.Duration
}

func (b *Barber) CutHair(customerID uuid.UUID) {
	log.Info("Cutting hair for Customer", zap.String("customerID", customerID.String()))
	sleep.RandomSleep(b.CutDuration)
	log.Info("Haircut done for Customer", zap.String("customerID", customerID.String()))
}
