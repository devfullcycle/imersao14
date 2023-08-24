// create unit test for route.go
package entity

import (
	"testing"
)

func TestNewRoute(t *testing.T) {
	route := NewRoute("1", "Route 1", 100)
	if route.ID != "1" {
		t.Errorf("Expected route ID to be '1', but got '%s'", route.ID)
	}
	if route.Name != "Route 1" {
		t.Errorf("Expected route Name to be 'Route 1', but got '%s'", route.Name)
	}
	if route.Distance != 100 {
		t.Errorf("Expected route Distance to be '100', but got '%f'", route.Distance)
	}
	if route.Status != "pending" {
		t.Errorf("Expected route Status to be 'pending', but got '%s'", route.Status)
	}
}

func TestFreightCalculate(t *testing.T) {
	route := NewRoute("1", "Route 1", 100)
	freight := NewFreight(0.5)
	freight.Calculate(route)
	if route.FreightPrice != 50 {
		t.Errorf("Expected route FreightPrice to be '50', but got '%f'", route.FreightPrice)
	}
}
