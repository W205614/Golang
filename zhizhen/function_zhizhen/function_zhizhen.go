// Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可
package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值: %d\n", a)
	fmt.Printf("交换前 b 的值: %d\n", b)

	swap(&a, &b)

	fmt.Printf("交换前 a 的值: %d\n", a)
	fmt.Printf("交换前 b 的值: %d\n", b)
}

func swap(x *int, y *int) {
	// var tmp int
	// tmp = *x
	// *x = *y
	// *y = tmp
	*x, *y = *y, *x
}
