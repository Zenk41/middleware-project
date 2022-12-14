package services

import (
	"middleware-project/models"
	"middleware-project/repository"
)

type AuthService struct {
	Repository repository.AuthRepository
}

func NewAuthService() AuthService {
	return AuthService {
		Repository: &repository.AuthRepo{},
	}
}

func (a *AuthService) Register(input models.UserInput) (*models.User, error) {
	return a.Repository.Register(input)
}

func (a *AuthService) Login(input models.UserLogin) string {
	return a.Repository.Login(input)
}


