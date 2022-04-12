package usecases

import (
	"go-base/pkg/crud/core/entities"
	"go-base/pkg/crud/core/ports"
	"go-base/pkg/helpers/configs"
)

type ExampleUseCase struct {
	ExampleRepositoryPort ports.ExampleRepositoryPort
}

func NewExampleUseCase(cf *configs.Config, exampleRepositoryPort ports.ExampleRepositoryPort) *ExampleUseCase {
	return &ExampleUseCase{exampleRepositoryPort}
}

func (u *ExampleUseCase) Create(example *entities.Example) error {
	err := u.ExampleRepositoryPort.Create(example)
	if err != nil {
		return err
	}
	return nil
}

func (u *ExampleUseCase) Read() ([]entities.Example, error) {
	data, err := u.ExampleRepositoryPort.Read()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *ExampleUseCase) Update(id string, example *entities.Example) error {
	err := u.ExampleRepositoryPort.Update(id, example)
	if err != nil {
		return err
	}
	return nil
}

func (u *ExampleUseCase) Delete(id string) error {
	err := u.ExampleRepositoryPort.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
