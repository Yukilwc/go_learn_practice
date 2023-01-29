package main

import "fmt"

// vscode断点调试方法：1 从当前文件run 2 指定固定文件run
// 如何配置lanuch文件

// 定义方式
// 内部方法声明
type Flyable interface {
	Fly() string
}

// 结构体定义
type Bird struct {
}

// 为结构体添加方法，此处得receiver如果是指针呢？如果是指针，则需要使用&Bird{}
func (b Bird) Fly() string {
	fmt.Println("Bird fly")
	return ""
}

func doFly(f Flyable) {
	f.Fly()
}

func main() {
	fmt.Println("start interface test")
	// 接口可被用来声明类型 零值为nil
	var a Flyable
	fmt.Println(a)
	// go实现接口，不需要显示声明
	var b Flyable = Bird{}
	b.Fly()
	// 其它函数，仅接收接口类型即可
	doFly((Bird{}))
	// 不使用implements关键字得好处
	// 业务类和接口送耦合，有了脚本语言的自由，同时还能有类型提示
	testAny()
	combineInterface()
}

// 空类型与类型断言
func testAny() {
	var anyVal interface{}
	var anyVal2 any
	fmt.Println("any val", anyVal, anyVal2)
	// 两种空类型定义都可以
	anyVal = anyVal2
	// 出现此情景，类似ts得any，一般后续需要类型断言来指定类型，才能继续使用
	// 使用 value.(Type)得方式，进行类型断言，并需要处理失败情景
	runWhenCanFly(Bird{})
	runWhenCanFly("I cannot fly")
	runWhenCanFlyUseSwitch(Bird{})
	runWhenCanFlyUseSwitch("I cannot fly")
	// 一种类似implements得编译阶段提示接口实现得写法,如果Bird遗漏了实现，编译器会提前报错
	var _ Flyable = Bird{}
}

// 测试类型断言
func runWhenCanFly(f any) {
	// 直接Fly会报错
	// f.Fly()
	if v, ok := f.(Flyable); ok {
		fmt.Println("f can fly,so ")
		v.Fly()
		return
	}
	fmt.Println("f cannot fly")
}

// 可以使用switch做一个多重判断类型
func runWhenCanFlyUseSwitch(f any) {
	switch f := f.(type) {
	case Bird:
		fmt.Println("switch case Flyable")
		f.Fly()
	case string:
		fmt.Println("I am a string", f)
	}
}

// 接口组合 可以把不同接口组合到新接口中 这在细粒度拆分接口时很有用
type Runner interface {
	Run() string
}
type FlyRun interface {
	Runner
	Flyable
}
type Chicken struct {
}

func (c Chicken) Fly() string {
	fmt.Println("Chicken fly")
	return ""
}
func (c Chicken) Run() string {
	fmt.Println("Chicken run")
	return ""
}
func combineInterface() {
	var a FlyRun = Chicken{}
	a.Fly()
	a.Run()
}
