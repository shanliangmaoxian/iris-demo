package controllers

import (
	"github.com/kataras/iris"
	"iris-demo/models"
	"iris-demo/service"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (c *UserController) GetAll() (result []models.User) {
	return c.Service.GetAll()
}
