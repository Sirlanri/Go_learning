package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)
func newApp() *iris.Application {
	app := iris.New()
	//可以从任何与http相关的panic中恢复，打印到控制台
	app.Use(recover.New())
	app.Use(logger.New())
	//控制器根路由路径 “/”
	mvc.New(app).Handle(new(ExampleController))
	return app
}
