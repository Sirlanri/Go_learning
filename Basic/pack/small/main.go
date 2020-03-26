package main

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/core/router"
)

func gogogo() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {})
	app.Run(iris.Addr(":8080"))
	app.PartyFunc("/cpanel", func(child iris.Party) {
		child.Get("/", func(ctx iris.Context) {})
	})
	// OR
	cpanel := app.Party("/cpanel")
	cpanel.Get("/", func(ctx iris.Context) {})
}

//敲敲路由
func handler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}
func myauthHandler(ctx iris.Context) {
	ctx.WriteString("认证失败")
}
func userProfileHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)

}
func userMsgHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
func router_learn() {
	app := iris.New()
	app.Get("/", handler)
	app.Post("/", handler)

	//party用于分组路由
	users := app.Party("/user", myauthHandler)
	users.Get("/{id:int}/profile", userProfileHandler)
	users.Get("/inbox/{id:int}", userMsgHandler)
}

func sonrouter() {
	app := iris.New()
	app.PartyFunc("/users", func(users iris.Party) {
		users.Use(mymiddle)
		users.Get("/{id:int}/porfile", userProfileHandler)
		users.Get("/inbox/{id:int}", userMsgHandler)
	})
}

//配套的小方法
func mymiddle(ctx iris.Context) {
	ctx.WriteString("认证失败")
	ctx.Next() //继续执行后续的handler
}
