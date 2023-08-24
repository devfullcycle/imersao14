package entity

import (
	"time"
)

type CustomTime time.Time

const layout = "2006-01-02T15:04:05.999Z"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	t, err := time.Parse(layout, s[1:len(s)-1]) // Remove quotes
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

type RouteRepository interface {
	Create(route *Route) error
	FindById(id string) (*Route, error)
	Update(route *Route) error
}

type Route struct {
	ID           string
	Name         string
	Distance     float64
	Status       string
	FreightPrice float64
	StartedAt    time.Time
	FinishedAt   time.Time
}

func NewRoute(id, name string, distance float64) *Route {
	return &Route{
		ID:       id,
		Name:     name,
		Distance: distance,
		Status:   "pending",
	}
}

func (r *Route) Start(startedAt time.Time) {
	r.Status = "started"
	r.StartedAt = startedAt
}

func (r *Route) Finish(finishedAt time.Time) {
	r.Status = "finished"
	r.FinishedAt = finishedAt
}

type FreightInterface interface {
	Calculate(route *Route)
}

type Freight struct {
	PricePerKm float64
}

func NewFreight(pricePerKm float64) *Freight {
	return &Freight{
		PricePerKm: pricePerKm,
	}
}

func (f *Freight) Calculate(route *Route) {
	route.FreightPrice = f.PricePerKm * route.Distance
}
