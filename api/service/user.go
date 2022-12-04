package service

import (
	"blogs/api/repository"
	"blogs/models"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserSerivce(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
