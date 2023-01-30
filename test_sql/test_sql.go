package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	initDB()
	// doQuery()
	// doQueryMulti()
	// name, err := inputName()
	// fmt.Println(" inputName name", name)
	// if err == nil {
	// 	doInsert(name)
	// }
	updateRow()
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

// 单行查询
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

// 多行查询
func doQueryMulti() {
	sqlStr := "select * from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {

	}
	defer rows.Close()
	var userList []User
	for rows.Next() {
		var u User
		rows.Scan(&u.id, &u.name)
		userList = append(userList, u)
	}
	fmt.Println("user list", userList)
}

// 插入
func doInsert(name string) {
	sqlStr := "insert into user(name) values (?)"
	ret, err := db.Exec(sqlStr, name)
	if err != nil {
		fmt.Println("insert failed", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed", err)
	}
	fmt.Println("insert success", theID)
}
func inputName() (string, error) {
	var name string
	fmt.Println("请输入名字:")
	n, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Errorf("scan in error:%s", err)
		return "", err
	}
	fmt.Println("scan in name", n)
	return name, nil
}

// 更新
func updateRow() {
	var id int
	var name string
	fmt.Println("请输入id和名字,用空格隔开，输入完成后回车:")
	_, err := fmt.Scanln(&id, &name)
	if err != nil {
		fmt.Println("scan in error", err)
		return
	}
	doUpdate(id, name)
}
func doUpdate(id int, name string) {
	sqlStr := "update user set name=? where id=?"
	ret, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Println("exec error", err)
		return
	}
	n, err := ret.RowsAffected()
	if err == nil {
		fmt.Println("rows n", n)
	}
}

func doDelete(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
