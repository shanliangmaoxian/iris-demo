package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"iris-demo/utils"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	config, _ := utils.GetYaml()

	app.Run(
		iris.Addr(":"+config.Server.Port),
		iris.WithCharset(config.Server.Charset),
		iris.WithoutServerError(iris.ErrServerClosed))
}
