package usecase

import (
	"fmt"

	"github.com/devfullcycle/imersao14/go/internal/routes/entity"
)

type CreateRouteInput struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Event    string  `json:"event"`
}

type CreateRouteOutput struct {
	ID           string
	Name         string
	Distance     float64
	Status       string
	FreightPrice float64
}

type CreateRouteUseCase struct {
	Repository entity.RouteRepository
	Freight    entity.FreightInterface
}

func NewCreateRouteUseCase(repository entity.RouteRepository, freight entity.FreightInterface) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repository: repository,
		Freight:    freight,
	}
}

func (u *CreateRouteUseCase) Execute(input CreateRouteInput) (*CreateRouteOutput, error) {
	route := entity.NewRoute(input.ID, input.Name, input.Distance)
	u.Freight.Calculate(route)
	fmt.Println(route)
	err := u.Repository.Create(route)
	if err != nil {
		return nil, err
	}
	return &CreateRouteOutput{
		ID:           route.ID,
		Name:         route.Name,
		Distance:     route.Distance,
		Status:       route.Status,
		FreightPrice: route.FreightPrice,
	}, nil
}
