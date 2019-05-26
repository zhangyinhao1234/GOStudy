package main

import (
	"fmt"
	"os"
	"runtime"
)
import "./com/zhangyinhao/test"

var VarTest_a string = "GO"

/**
变量的定义
*/
func main() {

	//test2()

	test3()

	//常用的变量

	vartest_1()
	vartest_2()
	vartest_3()

}

// init 函数在执行main之前加载
func init() {
	//init UserTest.init
	//开始执行本init
	VarTest_a = VarTest_a + " Study !!!  " + test.UserName
	fmt.Println(VarTest_a)
}

//step 1
func vartest_1() {
	fmt.Println(VarTest_a)
}

func vartest_2() {
	VarTest_a = "hello"
	fmt.Println(VarTest_a)
}

func vartest_3() {
	fmt.Println(VarTest_a)
}

func test3() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)

	path := os.Getenv("PATH")
	fmt.Printf("path is %s \n", path)
}

//定义变量
func test2() {
	var cc int
	var bb bool
	var dd string

	var (
		b1 int
		c1 bool
		d1 string
	)

	//快捷方式
	hh := 45
	cc = 15
	bb = false
	fmt.Println(hh)
	fmt.Println(cc)
	fmt.Println(bb)
	fmt.Println(dd)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

}
