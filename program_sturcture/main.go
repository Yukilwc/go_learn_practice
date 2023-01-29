package main

import (
	"fmt"
	testpkg "go_test_init/test_pkg"
)

func main() {
	// 赋值
	assiginment()
	testType()
	testPkg()
}
func assiginment() {
	a, b := 3, 4
	fmt.Println("元组赋值", a, b)
	numList := []string{"1", "2", "3"}
	fmt.Println("数组字面量初始化", numList)

}

func testType() {
	// type 别名 底层类型
	type Celsius float64
	type Fahrenheit float64
	const (
		a Celsius    = 2.33
		b Fahrenheit = 3.22
	)
	fmt.Println("类型别名", a, b)
	c := Celsius(b)
	fmt.Println("类型转换", c)
	// 类型断言得情景，更多是被迫定义为any，但是知道应该是某类型，此时断言
	// 而类型转换，是真的为了转换类型，应该比较少用
}

func testPkg() {
	// 导入时，根路径使用go.mod得模块声明
	// 最后得路径，是包上层得文件夹名字，而非包名，也不是包文件名
	// 路径前，可以对包进行别名命名
	// 包文件夹下得go文件，必须是同一个包名，否则会报错
	// import (
	// 	testpkg "go_test_init/test_pkg"
	// )
	fmt.Println("测试自定义包导入", testpkg.PkgName, testpkg.OtherPkgName)
}
