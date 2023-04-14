// go build -o local_timer.exe ./main.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("local_timer start")
	cron := cron.New()
	addGitPush(cron)
	cron.Start()
	defer cron.Stop()
	select {}
}
func addGitPush(cron *cron.Cron) {
	// spec := "0 0 18 * * *"
	spec := "19 19 * * *"
	_, err := cron.AddFunc(spec, doGitPush)
	if err != nil {
		fmt.Println(err)
	}
}
func doGitPush() {
	fmt.Println("do git push", time.Now().Format("2006-01-02 15:04:05"))
	pathList := []string{
		"D:\\workspace\\libiary\\ForTest\\go_code\\go_test_init\\",
		"D:\\workspace\\libiary\\my\\InclusiveLibrary\\",
		"D:\\workspace\\libiary\\my\\ProjectTemplate\\vue-docker\\",
		"D:\\workspace\\libiary\\my\\go-vue3-admin-template\\",
	}
	for _, v := range pathList {
		//切换到指定文件夹
		os.Chdir(v)
		//执行git命令
		cmd := exec.Command("git", "add", "-A")
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", "auto commit")
		cmd.Run()
		cmd = exec.Command("git", "push")
		cmd.Run()
		fmt.Println("push finished:", v)
	}

}
