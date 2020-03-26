package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	//添加两个内置处理程序 记录请求到终端
	app.Use(recover.New())
	app.Use(logger.New())

	//基于根路由的服务控制器
	mvc.New(app).Handle(new(ExampleController))

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	// http://localhost:8080/custom_path
	app.Run(iris.Addr(":8080"))
}

//ExampleController服务于 "/", "/ping" and "/hello".
type ExampleController struct{}

//从8080 Get
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>你好鸭</h1>",
	}
}

// http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "你好鸭混蛋！"}
}

// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "ping confirmed"
}

//  http://localhost:8080/custom_path
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func (ctx iris.Context){
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle("GET" "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
	return "hello from the custom handler without following the naming guide"
}
