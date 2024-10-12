package domain

import (
	"errors"
	"teste/internal/adapters/db/models"
	"teste/internal/ports"
)

type UserService struct {
	ports.UserRepository
	ports.ZipAdapter
}

func NewUserService(repository ports.UserRepository, zipAdapter ports.ZipAdapter) ports.UserService {
	return &UserService{UserRepository: repository, ZipAdapter: zipAdapter}
}

func (service *UserService) RegisterUser(user models.User, plainPassword string) (models.User, error) {
	if user.Name == "" || user.Email == "" || plainPassword == "" || user.Address.ZipCode == "" {
		return models.User{}, errors.New("nome, email, senha e CEP são obrigatórios")
	}

	address, err := service.ZipAdapter.GetAddressFromCorreiosByZipCode(user.Address.ZipCode)
	if err != nil {
		return models.User{}, err
	}

	user.Address.Street = address.Logradouro
	user.Address.Neighborhood = address.Bairro
	user.Address.City = address.Localidade
	user.Address.State = address.Uf

	user, err = service.UserRepository.Create(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (service *UserService) ListUsers() ([]models.User, error) {
	users, err := service.UserRepository.GetAll()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}
