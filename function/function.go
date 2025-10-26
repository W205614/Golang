package main

import "fmt"

// 默认情况下，Go 语言使用的是值传递
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func testmax() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值是: %d\n", ret)
}

// Go 函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

func testswap() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}

func main() {
	testmax()
	testswap()
}
