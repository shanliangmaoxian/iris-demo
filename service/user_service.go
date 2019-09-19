package service

import (
	"iris-demo/datasource"
	"iris-demo/models"
	"iris-demo/repo"
)

type UserService interface {
	GetAll() []models.User
}

type userService struct {
	repo *repo.UserDao
}

func NewUserService() *userService {
	return &userService{
		repo: repo.NewUserDao(datasource.InstanceMaster()),
	}
}

func (s *userService) GetAll() []models.User {
	return s.repo.GetAll()
}
