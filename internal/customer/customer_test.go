package customer

import (
	"github.com/google/uuid"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	if got == nil {
		t.Errorf("New() = nil, want non-nil")
	}
	if got.ID == uuid.Nil {
		t.Errorf("New().ID = %v, want non-nil UUID", got.ID)
	}
}
