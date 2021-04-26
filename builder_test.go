package designpatterns

import (
	"testing"
)

func TestHouseBuilder(t *testing.T) {
	builder := NewDirector(&CapeCod{})
	capecodHouse := builder.BuildHouse()
	if capecodHouse.Walls != "drywall" {
		t.Errorf("expected cape cod to have drywall, got %s", capecodHouse.Walls)
	}

	builder = NewDirector(&Igloo{})
	iglooHouse := builder.BuildHouse()
	if iglooHouse.Walls != "ice blocks" {
		t.Errorf("expected igloo to have ice blocks, got %s", iglooHouse.Walls)
	}
}
