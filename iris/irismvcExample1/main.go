package main

import (
	"mvctest/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main(){
	app := iris.New()
	app.Logger().SetLevel("debug")
	//加载模板文件
	app.RegisterView(iris.HTML("/web/views", ".html",)
	//注册控制器
	mvc.New(app.Party("/movies")).Handle(new(controllers.MovieController))
	app.Run(
		iris.Addr(
			iris.Addr("localhost:8090"),
			iris.WithOptimizations,
		);
}

func movies(app *mvc.Application){
	//添加身份验证的中间件
	app.Router.Use(middleware.BasicAuth)
	//使用内存数据创建资源库
	movieService := services.NewMovieService(repo)
	app.Register(movieService)
	app.Handle(new(controllers.MovieController))
}