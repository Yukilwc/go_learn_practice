package initialize

import (
	"gjm/config"
	"gjm/global"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接数据库
func InitDB() error {
	userName := config.DB["name"]
	password := config.DB["password"]
	path := "localhost"
	port := "3306"
	dbName := "ggg_test"
	mysqlConfig := mysql.Config{
		DSN: userName + ":" + password + "@tcp(" + path + ":" + port + ")/" + dbName + "?" + "charset=utf8&parseTime=true",
	}
	gormConfig := gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig); err != nil {
		return err
	} else {
		sqlDB, _ := db.DB()
		if err := sqlDB.Ping(); err != nil {
			return err
		} else {
			log.Println("数据库连接成功")
			global.DB = db
			return nil
		}
	}

}
