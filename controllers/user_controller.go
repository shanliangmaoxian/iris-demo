package controllers

import (
	"fmt"
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
	fmt.Println(err)
	if err != nil {
		log.Println(err)
	}
	eor := c.Service.Create(&info)
	if eor != nil {
		log.Println(eor)
	}
	return c.GetAll()
}

func (c *UserController) PostEdit() (result []models.User) {
	info := models.User{}
	err := c.Ctx.ReadForm(&info)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
	}
	eor := c.Service.Update(&info)
	if eor != nil {
		log.Println(eor)
	}
	return c.GetAll()
}

func (c *UserController) Delete() (result []models.User) {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return c.GetAll()
}

func (c *UserController) Get() *models.User {
	var data = &models.User{}
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		data = c.Service.Get(id)
	}
	return data
}
