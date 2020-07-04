package controllers

import (
	"mvc_login/datamodels"
	"mvc_login/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)
)

// UserController 是我们的/用户控制器。
// UserController 负责处理以下请求：
// GET              /user/register
// POST             /user/register
// GET                 /user/login
// POST             /user/login
// GET                 /user/me
//所有HTTP方法 /user/logout
type UserController struct {
	//每个请求都由Iris自动绑定上下文，
	//记住，每次传入请求时，iris每次都会创建一个新的UserController，
	//所以所有字段都是默认的请求范围，只能设置依赖注入
	//自定义字段，如服务，对所有请求都是相同的（静态绑定）
	//和依赖于当前上下文的会话（动态绑定）。
	Ctx iris.Context
	//我们的UserService，它是一个接口
	//从主应用程序绑定。
	Service services.UserService
	//Session，使用来自main.go的依赖注入绑定
	Session *sessions.Session
}

const userIDKey = "UserID"

func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *UserController) logout() {
	c.Session.Destroy()
}

var registerStaticView = mvc.View{
	Name:"user/register.html",
	Data: iris.Map{"title":"User Registration"},
}

// GetRegister 处理 GET: http://localhost:8080/user/register.
func (c *UserController) GetRegister() mvc.Result{
	//从表单中获取名字，用户名，密码
	var (
		firstname = c.Ctx.FormValue("firstname")
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)
	//创建新用户，密码哈希处理
	u, err := c.Service.Create(password,datamodels.User{
		Username: username,
		Firstname: firstname,
	})
	//讲用户的ID设置为此会话
	c.Session.Set(userIDKey,u.ID)
	return mvc.Response{
		//如果不是nil，就显示错误
		Err: err,
		Path: "/user/me",
		Code: 303,
	}
}

var loginStaticView=mvc.View{
	Name: "user/login.html",
	Data: iris.Map{"title":"User Login"},
}

//GetLogin GET: http://localhost:8080/user/login
func (c *UserController) GetLogin() mvc.Result{
	if c.isLoggedIn() {
		//如果已经登录了，就删除旧session
		c.loginout()
	}
	return loginStaticView
}

// PostLogin 处理POST: http://localhost:8080/user/register
func (c *UserController) PostLogin() mvc.Result{
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)
	u,found := c.Service.GetByUserNameAndPassword(username,password)
	if !found {
		return mvc.Response{
			Path: "/user/register",
		}
	}
	c.Session.Set(userIDKey,u.ID)
	return mvc.Response{
		Path: "/user/me",
	}
}

//GetMe 处理 GET: http://localhost:8080/user/me
func (c *UserController) GetMe() mvc.Result{
	if !c.isLoggedIn() {
		//没登录，就重定向到登录页面
		return mvc.Response{Path: "/user/login"}
	}
	u,found := c.Service.GetByID(c.getCurrentUserID())
	if !found {
		//如果session存在，但用户不在数据库中，注销
		c.logout()
		return c.GetMe()
	}
	return mvc.View{
		Name: "user/me.html",
		Data: iris.Map{
			"title":"资料是"+u.Username,
			"User":u,
		},
	}
}
// AnyLogout 处理 All/AnyHTTP 方法：http://localhost:8080/user/logout
func (c *UserController) AnyLogout(){
	if c.isLoggedIn() {
		c.logout()
	}
	c.Ctx.Redirect()
}

