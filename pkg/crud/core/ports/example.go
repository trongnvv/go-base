package ports

import "go-base/pkg/crud/core/entities"

type ExampleRepositoryPort interface {
	Create(example *entities.Example) error
	Read() ([]entities.Example, error)
	Update(id string, example *entities.Example) error
	Delete(id string) error
}
