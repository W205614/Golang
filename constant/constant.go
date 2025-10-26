// 语言常量
// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
// 常量的应用
package main

import "fmt"

func test1() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d", area)
	println()
	println(a, b, c)

	// 常量还可以用来枚举
	const (
		Unkown = 0
		Female = 1
		Male   = 2
	)
	println(Unkown, Female, Male)
}

func main() {
	test1()
}
