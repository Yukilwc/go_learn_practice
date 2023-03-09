package config

// 一般来说，config应该存储结构体，而结构体内容，应该存放在配置文件中
var DB = map[string]string{
	"name":     "root",
	"password": "123456abc",
}
