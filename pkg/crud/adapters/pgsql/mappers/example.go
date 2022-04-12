package mappers

import (
	"go-base/pkg/crud/adapters/pgsql/models"
	"go-base/pkg/crud/core/entities"
)

func ConverEntityToModel(e *entities.Example) *models.Example {
	return &models.Example{Name: e.Name}
}
