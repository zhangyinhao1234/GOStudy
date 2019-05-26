package main

import (
	"fmt"
	"strconv"
	"os"
)

//控制语句测试
func main() {

	var b bool = true
	if b{
		fmt.Println("true .....")
	}else{
		fmt.Println("false .....")
	}
	fmt.Printf("controllertest1 result:%s \n",controllertest1("go"))

	//初始化值处理
	var aa int =6
	if aa>10{
		fmt.Println("aa > 10")
	}
	//也可以这样写
	if bb :=6;bb<10{
		fmt.Println("bb < 10")
	}


	if ss := controllertest1("o");ss =="other" {
		fmt.Println("ss is not go ")
	}



	//swith 使用
	controllerswithtest(2)
	controllerswithtest(20)


	// d多返回值的用法


	s, bb := controllertest2("go")
	fmt.Printf("val is %s bool is %b",s,bb)


	ss2 :="ABC"
	ss2int, e := strconv.Atoi(ss2)
	if e !=nil{
		fmt.Printf("ss2 is not int %s \n",ss2)
		//如果想发生错误终止运行
		os.Exit(1)
		return
	}
	fmt.Printf(" value is %d \n",ss2int)


}

func controllertest1(hi string) string {
	if hi =="go"{
		return "go hello word"
	}else{
		return "other"
	}
}

func controllertest2(ss string) (ss2 string,bb bool)  {
	if ss =="hello"{
		return "hello",true
	}
	return  "go ",false
}

func controllerswithtest(i int)  {
	switch i {
	case 0:
		fmt.Println(".....0")
		return
	case 1,2,3:
		i+=9
		fmt.Println(i)
		return

	default:
		fmt.Println("default.....")
	}
}