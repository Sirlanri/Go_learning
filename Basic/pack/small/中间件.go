package middle

import "github.com/kataras/iris"

func middle_test() {
	//中间件测试（叫笔记更好点
	app := iris.New()
	app.Get("/", before, mainHandler, after)
	app.Run(iris.Addr(":8090"))
}
func before(ctx iris.Context) {
	shareInfor := "这是在中间件分享的信息"
	rePath := ctx.Path()
	println("在主Handler之前" + rePath)
	ctx.Values().Set("info", shareInfor)
	ctx.Next() //继续执行下一个handler，就是main
}
func after(ctx iris.Context) {
	println("主进程结束")
}
func mainHandler(ctx iris.Context) {
	println("现在是主进程了")
	//获取before中的信息
	info := ctx.Values().GetString("info")
	//回复客户端
	ctx.HTML("<h1>回复</h1>")
	ctx.HTML("<br/> Info: " + info)
	ctx.Next() //下一个，也就是after

}

func GlobalMiddle() {
	//全局使用中间件
	app := iris.New()
	//将“before”处理程序注册为将要执行的第一个处理程序
	//在所有域的路由上。
	//或使用`UseGlobal`注册一个将跨子域触发的中间件。
	app.Use(before)
	//after注册为将要执行的最后一个处理程序，在所有域路由处理程序之后
	app.Done(after)

	//注册路由
	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)
	app.Run(iris.Addr(":8080"))
}
