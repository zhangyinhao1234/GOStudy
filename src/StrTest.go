package main

import (
	"fmt"
	"strconv"
	. "strings"
)

//字符串测试
func main() {

	// \n 换行
	// \r 回车符
	// \t tab健
	// \u 或 \U：Unicode 字符
	// \\：反斜杠自身

	var ss string = "hello go"
	fmt.Printf("ss length is %d \n", len(ss))

	ss += " zhangyinhao"
	fmt.Printf("ss is is %s \n", ss)

	//strings.join 和 bytes.buffer效率会更好

	//判断字符串是否是 xxx开头
	prefix := HasPrefix(ss, "hello")
	fmt.Printf("prefix result is %t \n", prefix)

	//判断字符串是否是 xxx结尾
	suffix := HasSuffix(ss, "hello")
	fmt.Printf("suffix result is %t \n", suffix)

	fmt.Printf("ss contains zhang is %t \n", Contains(ss, "zhang"))

	fields := Fields(ss)
	fmt.Printf("fields is  %v \n", fields)
	for _, val := range fields {
		fmt.Printf("value is %s \n", val)
	}

	var tt = "zhang | huo|gogo"
	split := Split(tt, "|")
	for _, val := range split {
		fmt.Printf("split vallue is %s \n", val)
	}

	//strconv 对字符串进行转换

	var cc string = "100"
	i, _ := strconv.Atoi(cc)
	fmt.Printf("str conv count is %d \n", i)
	//转数字
	i += 5
	fmt.Printf("count to str is %s\n", strconv.Itoa(i))

	fmt.Printf("the size of int is %d \n", strconv.IntSize)

}
