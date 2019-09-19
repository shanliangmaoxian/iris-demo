package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-demo/controllers"
	"iris-demo/service"
)

func Configure(app *iris.Application) {
	userService := service.NewUserService()

	index := mvc.New(app.Party("/user"))
	index.Register(userService)
	index.Handle(new(controllers.UserController))
}
