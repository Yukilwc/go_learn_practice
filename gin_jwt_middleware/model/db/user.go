package db

import (
	"time"

	"gorm.io/gorm"
)

type CustomModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	UserName string `json:"userName" gorm:"index;comment:用户登录名"`
	Password string `json:"-"`
	Email    string `json:"email" gorm:"index"`
	NickName string `json:"nickName" gorm:"default:游客"`
	Enable   int    `json:"enable" gorm:"default:1;comment:是否启用"`
	CustomModel
}

func (User) TableName() string {
	return "user"
}
