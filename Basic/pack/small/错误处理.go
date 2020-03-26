package small

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func wrong() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notfound)
	app.OnErrorCode(iris.StatusInternalServerError, internal)
	app.Get("/", index)
	app.Run(iris.Addr(":8080"))
}
func notfound(ctx iris.Context) {
	ctx.View("errors/404.html")
}
func internal(ctx iris.Context) {
	ctx.WriteString("出错！")
}
func index(ctx context.Context) {
	ctx.View("index.html")
}

//MVC
func mvcTest() {
	app := iris.New()
	mvc.Configure(app.Party("/root"), myMvc)
	app.Run(iris.Addr(":8080"))
}
func myMvc(app *mvc.Application) {
	app.Handle(new(Mycontroller))
}

type Mycontroller struct{}

func (m *Mycontroller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/something/{id:long}", "MyCostomHandler", anyMiddleware...)
}

func (m *Mycontroller) Get() string {
	return "你好鸭"
}
func (m *myController) MyCustomHandler(id int64) string {
	return "Costomer Handler"
}
