package main

import (
	"fmt"
	"math"
)

// for　学习
func main() {

	//简单循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var ss1 string = ""
	for i := 0; i < 5; i++ {
		ss1 += "G"
		fmt.Printf("result: %s \n", ss1)
	}
	fmt.Printf("result: %s \n", ss1)

	math.Mod(float64(4), float64(2))

	//带条件的switch
	for i := 0; i < 100; i++ {
		mod := math.Mod(float64(i), float64(3))
		mod2 := math.Mod(float64(i), float64(5))
		switch {
		case mod == 0 && mod2 == 0:
			fmt.Println("FizzBuzz")
			break
		case mod == 0:
			fmt.Println("Fizz")
			break

		case mod2 == 5:
			fmt.Println("Buzz")
			break
		default:
			fmt.Println(i)

		}

	}

	k := 5
	for k >= 0 {
		k--
		fmt.Printf("k: %d \n", k)
	}

	//for range

	//遍历字符
	ss3 := "go hello zhang"
	for pos, char := range ss3 {
		//pos 索引位置
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

	for j := 0; j < 5; j++ {
		var vv int
		fmt.Printf("vv: %d \n", vv)
		vv = 5
	}

}
