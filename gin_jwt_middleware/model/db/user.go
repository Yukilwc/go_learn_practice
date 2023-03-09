package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName" gorm:"index;comment:用户登录名"`
	Password string `json:"-"`
	Email    string `json:"email" gorm:"index"`
	NickName string `json:"nickName" gorm:"default:游客"`
	Enable   int    `json:"enable" gorm:"default:1;comment:是否启用"`
}

func (User) TableName() string {
	return "user"
}
