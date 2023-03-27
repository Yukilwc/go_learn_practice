package initialize

import (
	"fmt"
	"gjm/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// 读取配置并存放到全局
func InitViper() {
	// 载入环境变量
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	env := os.Getenv("MISAKA_ENV")
	fmt.Println("当前环境变量", env)
	global.ENV = env
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file change", in.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			panic(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
}
