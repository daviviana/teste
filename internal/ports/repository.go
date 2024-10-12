package ports

import "teste/internal/adapters/db/models"

type Repository[T any] interface {
	Create(model T) (T, error)
	GetAll() ([]T, error)
}

type UserRepository interface {
	Repository[models.User]
	FindByEmail(email string) (models.User, error)
}
