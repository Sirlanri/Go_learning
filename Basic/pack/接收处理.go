package main

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/macro"
)

func leixing() {
	app := iris.New()
	app.Get("/username/{name}", func(ctx iris.Context) {
		fmt.Println("你好鸭，", ctx.Params().Get("name"))
	})
	//注册一个int类型的宏函数
	macro.Int.RegisterFunc("min", func(minValue int) func(string) bool {
		return func(papramValue string) bool {
			n, err := strconv.Atoi(papramValue)
			if err != nil {
				return false
			}
			return n >= minValue
		}
	})

	// http://localhost:8080/profile/id>=1
	app.Get("/profile/{id:int min(1)}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		fmt.Println("id是 ", id)
	})
	//改变路径参数错误引起的错误码
	app.Get("/profile/{id:int min(1)}/friends/{friendid:int min(1) else 504}",
		func(ctx iris.Context) {
			id, _ := ctx.Params().GetInt("id")
			friendid, _ := ctx.Params().GetInt("firendid")
			fmt.Println("ID是", id, friendid)
		})

}
