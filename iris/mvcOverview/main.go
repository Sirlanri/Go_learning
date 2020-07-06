package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//加载模板文件
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//服务控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))
	mvc.Configure(app.Party("/movies"), movies)

	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		iris.Addr(":8090"),
		//按下Ctrl+C跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//更快的json序列化/优化
		iris.WithOptimizations,
	)
}

func movies(app *mvc.Application) {
	//添加身份认证的中间件，基于movis请求
	app.Router.Use(middleware.BasicAuth)
	//使用数据源中的数据构建资料库
	repo := repositories.NewMovieRepository(datasource.Movies)
	//创建电影服务，绑定到电影应用程序的依赖项中
	movieService := services.NewMovieService(repo)
	app.Register(movieService)
	//为电影控制器提供服务
	app.Handle(new(controllers.MovieCtroller))
}
