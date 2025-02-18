package shop

import (
	"go-sleeping-barber/internal/customer"
	"go-sleeping-barber/internal/sleep"
	"go-sleeping-barber/internal/staff"
	logger "go-sleeping-barber/pkg/log"
	"go.uber.org/zap"
	"sync"
	"time"
)

var log = logger.GetLogger()

const DefaultCutDuration = 700 * time.Millisecond
const TimeToClose = 15 * time.Second

type BarberShop struct {
	Capacity         int
	WaitingCustomers chan customer.Customer
	BarberDone       chan bool
	Barber           staff.Barber
	wg               sync.WaitGroup
	Open             bool
}

func NewBarberShop(capacity int, barber staff.Barber) *BarberShop {
	return &BarberShop{
		Capacity:         capacity,
		WaitingCustomers: make(chan customer.Customer, capacity),
		BarberDone:       make(chan bool),
		Barber:           barber,
		Open:             true,
		wg:               sync.WaitGroup{},
	}
}

func (bs *BarberShop) StartBarber() {
	go func() {
		bs.Barber.Sleeping = false
		log.Info("Barber is ready to cut hair.")
		for {
			if len(bs.WaitingCustomers) == 0 {
				bs.Barber.Sleeping = true
				log.Info("No customers. Barber is sleeping.")
			}

			customerToCut, shopOpen := <-bs.WaitingCustomers
			if !shopOpen {
				log.Info("Shop is closed. Barber is done for the day!")
				bs.BarberDone <- true
				return
			}

			bs.Barber.Sleeping = false
			log.Info("Waking up to cut hair for Customer", zap.String("customerID", customerToCut.ID.String()))
			bs.Barber.CutHair(customerToCut)
		}
	}()
}

func (bs *BarberShop) addCustomer(customer customer.Customer) {
	if !bs.Open {
		log.Warn("Customer arrived after closing time", zap.String("customerID", customer.ID.String()))
		return
	}

	select {
	case bs.WaitingCustomers <- customer:
		log.Info("Customer arrived", zap.String("customerID", customer.ID.String()))
	default:
		log.Warn("Customer left because the shop is full", zap.String("customerID", customer.ID.String()))
	}
}

func (bs *BarberShop) close() {
	log.Info("Barber shop is closed.")
	bs.Open = false
	close(bs.WaitingCustomers)
	close(bs.BarberDone)
}

func Run() {
	shop := NewBarberShop(3, staff.Barber{CutDuration: DefaultCutDuration})
	shop.StartBarber()

	log.Info("Barber shop is open. Waiting for customers...")

	closed := make(chan bool)
	shopClosing := make(chan bool)

	go func() {
		for {
			select {
			case <-shopClosing:
				return
			case <-time.After(sleep.Random(100*time.Millisecond, 1*time.Second)):
				newCustomer := *customer.New()
				shop.addCustomer(newCustomer)
			}
		}
	}()

	go func() {
		<-time.After(TimeToClose)
		shopClosing <- true
		shop.close()
		closed <- true
	}()

	<-closed
}
