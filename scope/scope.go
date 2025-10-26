package main

import "fmt"

// 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量
func test1() {
	var a, b, c int

	a = 10
	b = 20
	c = a + b

	fmt.Printf("结果: a = %d, b = %d, c = %d\n", a, b, c)
}

// 在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用
var s int

func test2() {
	var a, b int

	a = 10
	b = 20
	s = a + b

	fmt.Printf("结果: a = %d, b = %d and s = %d\n", a, b, s)
}

// Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
var g int = 20

func test3() {
	var g int = 10

	fmt.Printf("结果: g = %d\n", g)
}

// 形式参数会作为函数的局部变量来使用
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}

func testsum() {
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main() 函数中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)
}

func main() {
	test1()
	test2()
	test3()
	testsum()
}
