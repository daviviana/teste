package repositories

import (
	"gorm.io/gorm"
)

type RepositoryGORM[T any] struct {
	db *gorm.DB
}

func NewRepositoryGORM[T any](db *gorm.DB) *RepositoryGORM[T] {
	return &RepositoryGORM[T]{db: db}
}

func (repository *RepositoryGORM[T]) Create(model T) (T, error) {
	err := repository.db.Create(&model).Error
	return model, err
}

func (repository *RepositoryGORM[T]) GetAll() ([]T, error) {
	var models []T
	result := repository.db.Find(&models)
	return models, result.Error
}
