package usecase

import (
	"testing"

	"github.com/devfullcycle/imersao14/go/internal/routes/entity"
)

// mock repository
func TestNewChangeRouteStatusUseCase(t *testing.T) {
	t.Run("should return a new ChangeRouteStatusUseCase", func(t *testing.T) {
		repository := new(MockRouteRepository)
		// return a new entity route
		repository.On("FindById", "1").Return(&entity.Route{
			ID:     "1",
			Name:   "Route 1",
			Status: "pending",
		}, nil)
		// update route
		repository.On("Update", &entity.Route{
			ID:     "1",
			Name:   "Route 1",
			Status: "started",
		}).Return(nil)
		useCase := NewChangeRouteStatusUseCase(repository)
		input := ChangeRouteStatusInput{
			ID:        "1",
			StartedAt: entity.CustomTime{},
			Event:     "RouteStarted",
		}
		output, err := useCase.Execute(input)
		if err != nil {
			t.Error(err)
		}
		if output.ID != input.ID {
			t.Error("ID is not equal")
		}
		if output.Status != "started" {
			t.Error("Status is not equal")
		}
	})
}
