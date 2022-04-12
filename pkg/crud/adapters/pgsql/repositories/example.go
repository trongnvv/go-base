package repositories

import (
	"go-base/pkg/crud/adapters/pgsql/mappers"
	"go-base/pkg/crud/adapters/pgsql/models"
	"go-base/pkg/crud/core/entities"
	"go-base/pkg/crud/core/ports"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ports.ExampleRepositoryPort {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) Create(example *entities.Example) error {
	e := mappers.ConverEntityToModel(example)
	result := r.db.Create(e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ExampleRepository) Read() ([]entities.Example, error) {
	var listExample []entities.Example
	result := r.db.Find(&listExample)
	if result.Error != nil {
		return nil, result.Error
	}
	return listExample, nil
}

func (r *ExampleRepository) Update(id string, example *entities.Example) error {
	e := mappers.ConverEntityToModel(example)
	result := r.db.Where("id = ", id).Updates(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ExampleRepository) Delete(id string) error {
	result := r.db.Delete(&models.Example{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
