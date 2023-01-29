package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	// testList()
	// testSlice()
	// testMap()
	// testStructure()
	testJson()
}

func testList() {
	// 数组由于固定尺寸，是比较少使用的
	// 数组长度是数组类型的一部分，[3]int和[4]int不是一个类型
	// 数组长度必须是常量表达式，是需要编译时确定的
	var a [3]int
	fmt.Println("数组初始化，需要指定长度", a)
	fmt.Println("下标独写值", a[0])
	// 循环数组
	for i, v := range a {
		fmt.Printf("%d--%d,", i, v)
	}
	var b [3]int = [3]int{1, 2, 3}
	fmt.Println("使用字面量进行初始化", b)
	// 字面量初始化语法糖 省去长度
	c := [...]int{1, 2}
	fmt.Println(c)
	l1 := [...]string{"a", "b"}
	l2 := [...]string{"a", "b"}
	fmt.Println("数组的比较是每一个值比较", l1 == l2)
	// 最后，数组是一种僵化的类型，一般都会使用切片这种动态化灵活的模式
}

func testSlice() {
	bottomArray := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println("容量为9得数组，类型[9]string", bottomArray)
	var s1 []string
	fmt.Println("常规初始化,默认零值为nil", s1, s1 == nil)
	s2 := bottomArray[1:3]

	fmt.Println("从一个底层数组上初始化,按惯例，都是从start开始，然后end不计入", s2)
	fmt.Println("start/end不写，默认最开始/结束", bottomArray[:3], bottomArray[1:], bottomArray[:])
	s3 := []int{1, 2, 3}
	fmt.Println("字面量初始化,直接一次建立底层数组+切片,数组和切片长度一致", s3)
	fmt.Println("使用make初始化make([]T,len,cap),会自动分配零值", make([]int, 10, 10))
	fmt.Println("基础信息获取,容量&尺寸", cap(s2), len(s2))
	// s4 := []int{1, 2, 3}
	fmt.Println("slice之间不可直接用==比较，数组可以,s3==s4会直接报错,需要手动for循环对比")
	fmt.Println("拼接元素", append(s3, 4, 5, 6))
	fmt.Println("拼接slice", append(s3, s3...))
}

func testMap() {
	var m1 map[string]int
	fmt.Println("map基础定义，零值为nil", m1, m1 == nil)
	// m1["alice"] = 1
	fmt.Println("直接向为分配空间的map存入元素会报panic，make后存入是可以的")
	m2 := make(map[string]int)
	fmt.Println("使用make分配空间", m2, m2 == nil, "make分配空间后，不再是nil")
	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("字面量初始化", m3)
	fmt.Println("长度也可以使用len获取", len(m3))
	m2["alice"] = 3
	fmt.Println("map不需要提前指定分配空间，make初始化后就可以用方括号语法赋值", m2, m2["alice"])
	delete(m2, "alice")
	fmt.Println("使用delete可以删除元素，并且是安全的", m2)
	for key, value := range m3 {
		fmt.Println("可以使用range for对map进行迭代", key, value)
	}
	fmt.Println("获取map不存在得key时，会返回一个value零值", m3["notDefinedKey"] == 0)
	v, ok := m3["notDefinedKey"]
	if !ok {
		fmt.Println("此key未定义")
	}
	fmt.Println("要区分零值是定义得，还是不存在导致得，可以使用第二个参数", v, ok)
	if v, ok := m3["notDefinedKey"]; !ok {
		fmt.Println("一个和if结合得语法糖，在if逻辑判断前先进行一次初始化", v)
	}
	fmt.Println("和slice一样，map不可直接使用==进行相等性比较，而是需要手动循环对比")
}

func testStructure() {
	type Employee struct {
		ID       int
		name     string
		Position string
		DoB      time.Time
	}
	fmt.Println("声明结构体")
	var a Employee
	fmt.Println("实例化结构体", a)
	a.name = "Misaka"
	fmt.Println("通过点号运算符访问", a)
	ap := &a
	fmt.Print("语法糖，结构体指针可以直接使用点号运算符操作，而不需要使用*获取值后再操作", ap.name, (*ap).name)
	type tree struct {
		value       int
		left, right *tree
	}
	var t tree
	fmt.Println("结构体不可包含自身，但是可包含自身指针", t)
	t2 := tree{value: 1}
	fmt.Println("字面值初始化,一般使用命名初始化", t2)
	fmt.Println("函数参数都是值拷贝传入，如果参数是大结构体，最好使用指针传入")
	tp2 := &tree{value: 2}
	fmt.Println("常用的初始化后取指针写法", tp2)
	c1 := tree{value: 3}
	c2 := tree{value: 3}
	fmt.Println("如果结构体全部成员可比较，那么结构体就是可比较的", c1 == c2)

	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	fmt.Println("以上，只声明成员类型，而不声明成员名的，被称为匿名成员,")
	var w Wheel
	w.X = 1
	fmt.Println("嵌入匿名成员时可直接访问叶子属性而不需要给出完整路径", Wheel{}, w, w.X)
	fmt.Println("嵌入匿名成员后，直接叶子级别的字面量初始化是不可以的，例如Wheel{X:1}是错误的")
	fmt.Println("也可以全路径访问,成员名字就是成员类型名", w.Circle.Point.X)
	fmt.Printf("格式化打印:%#v\n", w)

}

func testJson() {
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	movies := []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title:  "Bullitt",
			Year:   1968,
			Color:  true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshal failed: %s", err)
	}
	fmt.Printf("json data:%s\n", data)

	if data, err := json.MarshalIndent(movies, "", "    "); err == nil {
		fmt.Printf("带缩进打印:%s", data)
	}
	fmt.Println("tag中的omitempty表示为零值时，舍弃该字段")
	var ts []struct {
		Title  string
		others string
	}
	fmt.Println("可以只解码感兴趣的字段")
	if err := json.Unmarshal(data, &ts); err == nil {
		fmt.Printf("解码json:%#v\n", ts)
	}
}
