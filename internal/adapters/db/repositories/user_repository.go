package repositories

import (
	"teste/internal/adapters/db/models"
	"teste/internal/ports"
)

type UserRepository struct {
	*SQLiteRepository[models.User]
}

func NewUserRepository(databasePath string) ports.UserRepository {
	return &UserRepository{
		SQLiteRepository: NewSQLiteRepository[models.User](databasePath),
	}
}

func (userRepository *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := userRepository.db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
