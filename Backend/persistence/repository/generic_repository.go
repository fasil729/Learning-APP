package repositories

import (
	contracts "Brilliant/application/contracts/persistence"

	"gorm.io/gorm"
)

type GenericRepository[T any] struct {
	database *gorm.DB
}

func (repository *GenericRepository[T]) Create(data *T) (*T, error) {
	if err := repository.database.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repository *GenericRepository[T]) Delete(data *T) (*T, error) {
	if err := repository.database.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repository *GenericRepository[T]) GetAll() ([]*T, error) {
	var result []*T
	if err := repository.database.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *GenericRepository[T]) GetById(id uint) (*T, error) {
	var result T
	if err := repository.database.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (repository *GenericRepository[T]) Update(data *T) (*T, error) {
	if err := repository.database.Save(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewGenericRepository[T any](getDb func() *gorm.DB) contracts.IGenericRepository[T] {
	return &GenericRepository[T]{
		database: getDb(),
	}
}
