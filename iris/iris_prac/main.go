package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/testtwo", testtwo)
	app.Get("/{id:int}/{name:string}", getname)

	app.Post("/postone", postone)
	//子路径出了点问题....
	test := app.Party("/testone", testoneHand)
	test.Get("/{id:int}", getid)
	test.Get("/{id:int}/{name:string}", getname)
	app.Run(iris.Addr("localhost:8090"))
}

func testoneHand(ctx iris.Context) {
	ctx.Text("第一个请求成功啦~")
}
func getid(ctx iris.Context) {
	//获取ID
	ctx.Text("你的ID是" + ctx.Params().Get("id"))
}
func getname(ctx iris.Context) {
	//返回ID+名字，以json格式
	surprise := User{
		Idnum: ctx.Params().Get("id"),
		Name:  ctx.Params().Get("name"),
		Other: "这是用GET得到的第一条json哦~",
	}
	ctx.JSON(surprise)
}
func testtwo(ctx iris.Context) {
	ctx.Text("通过主路径返回的")
}
func postone(ctx iris.Context) {
	var user User

	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(301) //返回给前端的错误代码
		fmt.Println(err)
	} else {
		fmt.Println("获取到json！")
		fmt.Println(user.Idnum, user.Name, user.Other)
		ctx.Text("成功接收了你的参数")
	}

}

//用于返回json的数据格式
type User struct {
	Idnum string `json:"idnum"`
	Name  string `json:"name"`
	Other string `json:"other"`
}
