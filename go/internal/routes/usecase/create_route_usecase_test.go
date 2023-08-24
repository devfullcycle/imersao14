// create usecase test using testify mocks
package usecase

import (
	"testing"

	"github.com/devfullcycle/imersao14/go/internal/routes/entity"
	"github.com/stretchr/testify/mock"
)

type MockRouteRepository struct {
	mock.Mock
}

func (m *MockRouteRepository) Create(route *entity.Route) error {
	args := m.Called(route)
	return args.Error(0)
}

func (m *MockRouteRepository) FindById(id string) (*entity.Route, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Route), args.Error(1)
}

func (m *MockRouteRepository) Update(route *entity.Route) error {
	args := m.Called(route)
	return args.Error(0)
}

func TestCreateRouteUseCaseExecute(t *testing.T) {
	repository := new(MockRouteRepository)
	freight := entity.NewFreight(0.5)
	useCase := NewCreateRouteUseCase(repository, freight)
	input := CreateRouteInput{
		ID:       "1",
		Name:     "Route 1",
		Distance: 100,
	}
	route := entity.NewRoute(input.ID, input.Name, input.Distance)
	repository.On("Create", route).Return(nil)
	freight.Calculate(route)
	output, err := useCase.Execute(input)
	if err != nil {
		t.Error(err)
	}
	repository.AssertExpectations(t)
	if output.ID != input.ID {
		t.Error("ID is not equal")
	}
	if output.Name != input.Name {
		t.Error("Name is not equal")
	}
	if output.Distance != input.Distance {
		t.Error("Distance is not equal")
	}
	if output.Status != route.Status {
		t.Error("Status is not equal")
	}
	if output.FreightPrice != route.FreightPrice {
		t.Error("FreightPrice is not equal")
	}
}
