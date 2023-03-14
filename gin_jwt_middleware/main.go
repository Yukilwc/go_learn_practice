package main

import (
	"fmt"
	"gjm/global"
	"gjm/initialize"
)

func main() {
	fmt.Println("项目启动")
	initialize.InitViper()
	fmt.Println("配置读取完成:", global.CONFIG)
	if err := initialize.InitDB(); err != nil {
		// log.Fatal(err.Error())
		panic(err)
	}
	initialize.InitRouter()

}
