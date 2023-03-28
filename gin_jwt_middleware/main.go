package main

import (
	"fmt"
	"gjm/initialize"
)

func main() {
	fmt.Println("项目启动")
	initialize.InitViper()
	initialize.InitLogger()
	// global.LOG.Info("配置读取完成:", zap.Any("config", global.CONFIG))
	if err := initialize.InitDB(); err != nil {
		// log.Fatal(err.Error())
		panic(err)
	}
	initialize.StartTimer()
	initialize.InitRouter()

}
