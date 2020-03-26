package main

import (
	"database/sql"
	"fmt"

	//下面这玩意儿还必须得引用啊
	_ "github.com/go-sql-driver/mysql"
)

func main1() {
	db, err := sql.Open("mysql", "root:123456@/gotest")
	if err != nil {
		fmt.Println("连接失败", err)
	}

	//selectTest(db)
	//insertTest(db, "再来一次吧")
	//updateTest(db, "自动升级的文字")
	yubianyi1(db)
}

func updateTest(db *sql.DB, newtext string) {
	result, err := db.Exec("update goauto set names=? where id=3", newtext)
	if err != nil {
		fmt.Println("更新出错", err)
	} else {
		rownum, _ := result.RowsAffected()
		fmt.Println("影响的行数", rownum)
	}
}
func selectTest(db *sql.DB) {
	aveScore := 500
	rows, err2 := db.Query("select 专业名称,录取最低分 from 17lg where 平均分>= ?", aveScore)
	if err2 != nil {
		fmt.Println("执行SQL出错：", err2)
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		var minscore int
		if err := rows.Scan(&name, &minscore); err != nil {
			fmt.Println("读取row出错", err)
		}
		fmt.Println(name, minscore)
	}
}
func insertTest(db *sql.DB, name string) {
	result, err := db.Exec("insert into goauto values(null,?)", name)
	if err != nil {
		fmt.Println("插入出错", err)
	}
	if lastid, err := result.LastInsertId(); err == nil {
		fmt.Println("本次插入的ID是", lastid)
	}
	if rowsEffect, err := result.RowsAffected(); err == nil {
		fmt.Println("影响的行数：", rowsEffect)
	}
}

//采用预编译方法写的sql语句
func yubianyi1(db *sql.DB) {
	stmt, _ := db.Prepare("insert into goauto values(null,?)")

	defer stmt.Close()

	result, err := stmt.Exec("预编译，第一次插入")
	if err != nil {
		fmt.Println("预编译插入数据错误：", err)
	}
	lastid, _ := result.LastInsertId()
	affectRows, _ := result.RowsAffected()
	fmt.Println("这次插入的ID是：", lastid)
	fmt.Println("影响的行数：", affectRows)
}
