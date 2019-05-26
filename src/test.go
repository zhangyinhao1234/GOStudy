package main

import "fmt"

//引用其他包文件
import "./com/zhangyinhao/test"

const a = "hello"

var c int = 5
var r int = 5

type T struct {
}

func main() {

	var v int
	hi()
	fmt.Println(v)

	//引用其他包文件和变量
	userName := test.UserName
	fmt.Println(userName)

}

func hi() {

}

func test1(t T) {

}
