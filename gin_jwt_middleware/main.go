package main

import (
	"fmt"
	"gjm/initialize"
)

func main() {
	fmt.Println("Gin jwt 中间件")
	initialize.InitRouter()
}
