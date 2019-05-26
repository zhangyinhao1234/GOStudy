package main

import "fmt"

/**
常量的学习和测试
*/

const pi = 3.1415926
const userName = "yinhao"
const age int = 29
const c1 = 2 / 3

const Ln2 = 123456

//一次定义好几个常量
const f1, f2, f3 = 1, "go", " hello"
const Monday, Tuesday = 1, 2

//作为枚举
const (
	Monday_1  = 1
	Tuesday_1 = 2
)

// iota 的使用
const (
	Sunday    = iota //0
	Monday_2         //1
	Tuesday_2        //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

//type的使用
//作为别名
type Color int

const (
	RED    Color = 1
	BIUE   Color = 2
	YELLOW Color = 3
)

func main() {
	fmt.Println(testInt())
	//fmt.Print(c1)
	fmt.Println(Saturday)
}

func constTest_1() {
	fmt.Printf("constTest_1 %s", pi)
}

//测试数字转换
func testInt() int {
	return age + 5
}
