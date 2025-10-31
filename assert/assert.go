package main

import "fmt"

func test1() {
	var i interface{} = "Hello, Go!"

	// 尝试将 i 断言为 string 类型
	s, ok := i.(string)
	if ok {
		fmt.Println("断言成功:", s)
	} else {
		fmt.Println("断言失败")
	}

	// 尝试将 i 断言为 int 类型
	n, ok := i.(int)
	if ok {
		fmt.Println("断言成功:", n)
	} else {
		fmt.Println("断言失败")
	}
}

func test2() {
	var i interface{} = "Hello, Go!"

	// 直接断言为 string 类型
	s := i.(string)
	fmt.Println("断言成功:", s)

	// 直接断言为 int 类型（会引发 panic）
	n := i.(int)
	fmt.Println("断言成功:", n)
}

func main() {
	test1()
	test2()
}
