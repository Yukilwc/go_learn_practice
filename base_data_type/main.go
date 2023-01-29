package main

import "fmt"

func main() {
	// 基础数据类型
	testNum()
	boolTest()
}
func testNum() {
	var a int = 3
	var b int32 = 4
	var c float32 = 1.1
	var d float64 = 2.2

	fmt.Printf("整数:%d,%d,浮点数:%g,%g", a, b, c, d)
}

func boolTest() {
	// And和Or运算符，且存在
	fmt.Println(true && false, true || false)
}
