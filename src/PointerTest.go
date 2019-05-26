package main

import "fmt"

//指针的测试
func main() {

	var aa int = 5
	fmt.Printf("aa 数据：%d 地址： %p \n", aa, &aa)

	//指针定义
	var bb *int
	bb = &aa
	fmt.Printf("指针地址 %p \n", bb)

	//给指针设置数据。 aa也就变了
	*bb = 10
	fmt.Printf("bb id %d,aa id %d \n", *bb, aa)

}
