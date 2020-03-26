package main

import (
	"database/sql"
	"fmt"
)

//把用户信息写入数据库
func Adduser(user *User) bool {
	db, _ := sql.Open("mysql", "root:12356@tcp(127.0.0.1:3306)/gotest")

	result, err1 := db.Exec("insert into webtest(idnum,name,other)values(?,?,?)", user.Idnum, user.Name, user.Other)
	if err1 != nil {
		fmt.Println("插入出错啦！", err1.Error())
		return false
	}
	lastInsertID, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println("获取ID出错", err2.Error())
		return false
	}
	fmt.Println("插入的ID是", lastInsertID)

	rowsAffect, err2 := result.RowsAffected()
	fmt.Println("影响了", rowsAffect)
	return true
}

//获取所有人，返回一个user列表
func GetAllusers() []*User {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
	single := new(User)
	rows, err := db.Query("select idnum,name,other form webtest where id = ?",)
	defer func(){
		if 
	}
}
