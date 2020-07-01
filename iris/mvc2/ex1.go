package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main1() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	mvc.Configure(app.Party("basic"), basicMVC)
	app.Run(iris.Addr(":8090"))
}

func basicMVC(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("路径是%s", ctx.Path())
		ctx.Next()
	})
	//把依赖注入controler绑定
	app.Register(
		sessions.New(sessions.Config{}).Start,
		&prefixedLogger{prefix: "DEV"},
	)
	app.Handle(new(basicController))
	app.Party("/sub").Handle(new(basicSubController))
}

type LoggerService interface {
	Log(string)
}

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s:%s", s.prefix, msg)
}

type basicController struct {
	Logger  LoggerService
	Session *sessions.Session
}

func (c *basicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/custom", "Custom")
}

func (c *basicController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController 报错啦")
	}
}

func (c *basicController) Get() string {
	count := c.Session.Increment("count", 1)
	body := fmt.Sprintf("Hello from basicController\nTotal visits from you: %d", count)
	c.Logger.Log(body)
	return body
}

func (c *basicController) Custom() string {
	return "custom"
}

type basicSubController struct {
	Session *sessions.Session
}

func (c *basicSubController) Get() string {
	count := c.Session.GetIntDefault("count", 1)
	return fmt.Sprintf("Hello from basicSubController.\nRead-only visits count: %d", count)
}
