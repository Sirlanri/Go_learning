package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/gotest?charset=utf8")

	if err != nil {
		panic(err)
	}
	err = db.Ping() //测试能否链接
	if err != nil {
		fmt.Println(err)
	}

	Get(db)
	Gets(db)
	Insert(db)
}

func Get(db *sql.DB) {
	var (
		id       int
		username string
		number   int
	)
	query := `SELECT id, name, num FROM first WHERE id=?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, username, number, "\n单行读取完成")
}

func Gets(db *sql.DB) {
	type user struct {
		id       int
		username string
		number   int
	}
	rows, _ := db.Query(`SELECT id, name, num FROM first`)
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.number)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, u)
	}
	err := rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range users { //前面默认会带个key，不能省略
		fmt.Println(item.id, item.username, item.number)

	}
}

func Insert(db *sql.DB) {
	result, err := db.Exec(`INSERT INTO first (id,name,num) VALUES(?,?,?)`, 4, "新四", 130)
	if err != nil {
		println(result, err)
	}
	fmt.Println("插入完毕")

}
