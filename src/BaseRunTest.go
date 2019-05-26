package main

import (
	"fmt"
	"math/rand"
	"time"
)

//基本类型和运算符测试
func main() {

	var b bool = true
	fmt.Printf("bool: %s", b)
	var cc = 10
	var dd = "10"
	fmt.Printf("bool: true : %S", cc != 5)
	fmt.Printf("bool: false: %s", cc == 5)

	//编译异常 只有两个类型相同的值才可以进行比较
	//fmt.Printf("cc == dd result: %S",cc == dd)

	//运算符   !（非） &&(与) ||（或）
	var ff bool = true
	var gg bool = false
	fmt.Printf("str : 10 and 11 : %s", dd != "11")
	fmt.Printf("ff ! %S:", !ff)
	fmt.Printf("ff && gg %S:", gg && ff)
	fmt.Printf("ff || gg %S \n", gg || ff)

	//浮点型
	//float32 精确到小数点7位
	f := float32(3.141592653)
	//float64 精确到小数点15位 尽可能的使用float64
	i := float64(3.141592653)
	fmt.Printf("float32: %g \n", f)
	fmt.Printf("float64: %g \n", i)

	//整形
	//int8（-128 -> 127）
	//int16（-32768 -> 32767）
	//int32（-2,147,483,648 -> 2,147,483,647）
	//int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
	var boo int
	var foo int32
	boo = 15

	//编译异常,强类型所以编译异常
	//foo = boo + boo

	fmt.Println(boo)
	foo = foo + 10
	fmt.Printf("foo: %d \n", foo)

	//常用函数

	for n := 0; n < 10; n++ {
		i2 := rand.Int()
		fmt.Printf("rand : %d \n", i2)
	}

	for m := 0; m < 5; m++ {
		// 返回 [0,n)之间的伪随机数
		intn := rand.Intn(100)
		fmt.Printf("rand.Intn : %d \n", intn)
	}

	nanosecond := time.Now().Nanosecond()
	fmt.Printf("timestamp : %d \n", nanosecond)

	// 类型别名 type
	type Count int
	var f11, f22, f33 Count = 1, 2, 3
	f44 := f22 + f33 + f11
	fmt.Print(f44)
	fmt.Printf("\n")

	//字符类型  byte
	var ch byte = 'A'
	var ch2 byte = 65
	fmt.Print(ch)
	fmt.Printf("\n")
	fmt.Print(ch2)

}
