package initialize

import (
	"gjm/global"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接数据库
func InitDB() error {
	mysqlConfig := mysql.Config{
		DSN: global.CONFIG.Mysql.Dsn(),
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
