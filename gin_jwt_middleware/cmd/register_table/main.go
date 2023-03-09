package main

// 表创建工具
import (
	"gjm/global"
	"gjm/initialize"
	"gjm/model/db"
	"log"
)

// 创建数据库表

func main() {
	if err := initialize.InitDB(); err != nil {
		// log.Fatal(err.Error())
		panic(err)
	}
	err := global.DB.AutoMigrate(&db.User{})
	if err != nil {
		panic(err)
	} else {
		log.Println("创建数据库表成功")
	}
}
