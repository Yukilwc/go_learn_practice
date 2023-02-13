package main

// 编译指令
// go build -o ./minappsync/minappsync.exe  ./minappsync/main.go

import (
	"fmt"
)

func main() {
	fmt.Println("小程序同步Start")
	// reader := bufio.NewReader(os.Stdin)
	// text, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("text:", text)
	// fmt.Scanln(&text)
	// fmt.Println("text scanln:", text)
	// fmt.Scanln(&text)
	// fmt.Println("text scanln:", text)
	loop()
}

var mpFolder = "D:/workspace/work/web/gitee/etranscode/etransmp3IM"
var backupFolder = "D:/workspace/work/web/gitee/etranscode/etransmp3Backup"

func loop() {
	var text string
	for {
		fmt.Println("请输入你的转移目标(mp/backup):", text)
		fmt.Scanln(&text)
		if text == "exit" || text == "quit" {
			break
		}
	}
}

func mp2backup() {

}

func backup2mp() {

}
