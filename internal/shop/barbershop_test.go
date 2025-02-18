package shop

import (
	"go-sleeping-barber/internal/customer"
	"go-sleeping-barber/internal/staff"
	"testing"
)

func TestNewBarberShop(t *testing.T) {
	barber := staff.Barber{CutDuration: DefaultCutDuration}

	shop := NewBarberShop(3, barber)

	if shop.Capacity != 3 {
		t.Errorf("Expected capacity 3, got %d", shop.Capacity)
	}
	if shop.Barber.CutDuration != DefaultCutDuration {
		t.Errorf("Expected cut duration %v, got %v", DefaultCutDuration, shop.Barber.CutDuration)
	}
	if !shop.Open {
		t.Errorf("Expected shop to be open")
	}
}

func TestAddCustomer(t *testing.T) {
	barber := staff.Barber{CutDuration: DefaultCutDuration}
	shop := NewBarberShop(3, barber)
	client := customer.New()

	shop.addCustomer(*client)

	if len(shop.WaitingCustomers) != 1 {
		t.Errorf("Expected 1 client, got %d", len(shop.WaitingCustomers))
	}
}

func TestCloseShop(t *testing.T) {
	barber := staff.Barber{CutDuration: DefaultCutDuration}
	shop := NewBarberShop(3, barber)

	shop.close()

	if shop.Open {
		t.Errorf("Expected shop to be closed")
	}
}
