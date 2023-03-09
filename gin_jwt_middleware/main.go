package main

import (
	"fmt"
	"gjm/initialize"
)

func main() {
	fmt.Println("Gin jwt 中间件")
	if err := initialize.InitDB(); err != nil {
		// log.Fatal(err.Error())
		panic(err)
	}
	initialize.InitRouter()

}
