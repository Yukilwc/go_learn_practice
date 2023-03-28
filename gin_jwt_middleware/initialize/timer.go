package initialize

import (
	"fmt"
	"gjm/global"
	"time"

	"go.uber.org/zap"
)

// 启动定时

func StartTimer() {
	id, err := global.Timer.AddTaskByFunc("testTimer", "*/1 * * * *", func() {
		fmt.Println("定时打印:", time.Now().Format("2006-01-02 03:04:05"))
	})
	if err != nil {
		global.LOG.Error("定时启动失败:" + err.Error())
		return
	}
	global.LOG.Info("定时启动成功:", zap.Any("id", id))
}
