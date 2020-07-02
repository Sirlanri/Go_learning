package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
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

func main() {
	app := newApp()
	app.Run(iris.Addr(":8090"))
}

//ExampleController 提供 ”/”，“/ping”和 “/hello”路由选项
type ExampleController struct{}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>你好</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong~~~"
}

func (c *ExampleController) GetHello() interface{} {
	return map[string]string{
		"message": "hello",
	}
}

//main调用controller之前，先调用Before
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddleware := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("正在执行中间件custompath")
		ctx.Next()
	}
	//不用命名规则的自定义handler
	b.Handle("GET", "/custompath", "thishandle", anyMiddleware)
}

func (c *ExampleController) thishandle() string {
	return "从custom handle的消息，没有遵守Get命名规则哦"
}
