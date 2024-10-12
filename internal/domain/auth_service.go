package domain

import (
	"errors"
	"teste/internal/ports"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("mysecret") // Chave secreta para JWT

type AuthService struct {
	ports.UserRepository
	ports.EmailAdapter
}

func NewLoginService(repository ports.UserRepository, adapter ports.EmailAdapter) ports.AuthService {
	return &AuthService{UserRepository: repository, EmailAdapter: adapter}
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	if !user.CheckPassword(password) {
		return "", errors.New("credenciais inválidas")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) GenerateRecoveryToken(email string) (string, error) {
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	service.EmailAdapter.SendEmail()

	return tokenString, nil
}
