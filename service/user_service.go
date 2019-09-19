package service

import (
	"iris-demo/datasource"
	"iris-demo/models"
	"iris-demo/repo"
)

type UserService interface {
	GetAll() []models.User
	Get(id int) *models.User
	Delete(id int) error
	Update(user *models.User) error
	Create(user *models.User) error
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

func (s *userService) Create(data *models.User) error {
	return s.repo.Create(data)
}

func (s *userService) Update(data *models.User) error {
	return s.repo.Update(data)
}

func (s *userService) Get(id int) *models.User {
	return s.repo.Get(id)
}

func (s *userService) Delete(id int) error {
	return s.repo.Delete(id)
}
