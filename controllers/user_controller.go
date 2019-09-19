package controllers

import (
	"github.com/kataras/iris"
	"iris-demo/models"
	"iris-demo/service"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (c *UserController) GetAll() (result []models.User) {
	return c.Service.GetAll()
}

func (c *UserController) PostSave() (result []models.User) {
	info := models.User{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}
	c.Service.Create(&info)
	return c.GetAll()
}
