package ports

import "teste/internal/adapters/db/models"

type UserService interface {
	RegisterUser(user models.User, plainPassword string) (models.User, error)
	ListUsers() ([]models.User, error)
}

type AuthService interface {
	Login(email, password string) (string, error)
	GenerateRecoveryToken(email string) (string, error)
}
