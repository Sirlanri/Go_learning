package controllers

import (
	"errors"
	"mvctest/datamodels"
	"mvctest/services"

	"github.com/kataras/iris"
)

type MovieController struct {
	//是一个界面，从主程序绑定
	Service services.MovieService
}

// 获取电影列表
// curl -i http://localhost:8080/movies
func (c *MovieController) Get() (results []datamodels.Movie) {
	return c.Service.Getall() //返回数据类型？
}

//获取一部电影
// curl -i http://localhost:8080/movies/1
func (c *MovieController) GetBy(id int64) (movie datamodels.Movie, found bool) {
	return c.Service.GetByID(id) //如果没找到，就报404
}

//用put来获取一部（我学了一个月前端还从没用过put
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MovieController) PutBy(ctx iris.Context, id int64) (datamodels.Movie, error) {
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels.Movie{}, errors.New("获取失败啦！")
	}
	file.Close()
	//这是上传文件的网址
	poster := info.Filename
	genre := ctx.FormValue("genre")
	return c.Service.UpdataPosterAndGenreByID(id, poster, genre)
}

//用delete删除电影
//curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *MovieController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		//返回删除的id
		return iris.StatusBadRequest
	}
	return iris.StatusBadRequest
}
