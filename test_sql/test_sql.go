package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	initDB()
	doQuery()
	defer db.Close()
}
func initDB() {
	dsn := "root:123456abc@tcp(127.0.0.1:3306)/my_test_db?charset=utf8mb4&parseTime=True"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping err", err)
	}
}

type User struct {
	id   int
	name string
}

func doQuery() {
	var u User
	sqlStr := "select * from user where id=?"
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name)
	if err != nil {
		fmt.Println("query error", err)
		return
	}
	fmt.Printf("id:%d name:%s\n", u.id, u.name)
}
