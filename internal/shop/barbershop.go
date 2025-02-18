package shop

import (
	"go-sleeping-barber/internal/sleep"
	"go-sleeping-barber/internal/staff"
	logger "go-sleeping-barber/pkg/log"
	"go.uber.org/zap"
	"sync"
	"time"

	"github.com/google/uuid"
)

var log = logger.GetLogger()

const DefaultCutDuration = 500 * time.Millisecond

type BarberShop struct {
	Capacity         int
	WaitingCustomers chan uuid.UUID
	BarberDone       chan bool
	Barber           staff.Barber
	wg               sync.WaitGroup
	Open             bool
}

func NewBarberShop(capacity int) *BarberShop {
	return &BarberShop{
		Capacity:         capacity,
		WaitingCustomers: make(chan uuid.UUID, capacity),
		BarberDone:       make(chan bool),
		Barber:           staff.Barber{Sleeping: false, CutDuration: DefaultCutDuration},
		Open:             true,
	}
}

func (bs *BarberShop) StartBarber() {
	go func() {
		defer bs.wg.Done()

		for {

			if len(bs.WaitingCustomers) == 0 {
				bs.Barber.Sleeping = true
				log.Info("No customers. Barber is sleeping.")
			}

			customerID, ok := <-bs.WaitingCustomers
			if !ok {
				log.Info("Shop is closed. Barber is done for the day!")
				bs.BarberDone <- true
				return
			}

			bs.Barber.Sleeping = false
			log.Info("Waking up to cut hair for Customer", zap.String("customerID", customerID.String()))
			bs.Barber.CutHair(customerID)
			bs.BarberDone <- true
		}
	}()
}

func Run() {
	shop := NewBarberShop(4)
	shop.StartBarber()

	log.Info("Barber shop is open. Waiting for customers...")

	shop.Open = false
	close(shop.WaitingCustomers)

	sleep.RandomSleep(2 * time.Second)

	log.Info("Barber shop is closed.")
}
