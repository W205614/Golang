package main

import "fmt"

// if 语句
func test_if() {
	var a int = 10

	if a < 20 {
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为: %d\n", a)
}

// if_else 语句
func test_if_else() {
	var a int = 100

	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为: %d\n", a)
}

// if 嵌套
func test_ifif() {
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为 100, b 的值为 200\n")
		}
	}
	fmt.Printf("a 的值为: %d\n", a)
	fmt.Printf("b 的值为: %d\n", b)
}

func main() {
	test_if()
	test_if_else()
	test_ifif()
}
