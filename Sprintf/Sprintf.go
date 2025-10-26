package main

import (
	"fmt"
	"io"
	"os"
)

// Sprintf 根据格式化参数生成格式化的字符串并返回该字符串
func test1() {
	var stockcode = 123
	var enddate = "2025-10-23"
	var url = "Code = %d &enddate = %s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s)
}

func main() {
	test1()
}
