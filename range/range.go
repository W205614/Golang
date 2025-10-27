// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 数组和切片
func test1() {
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}

// 字符串
// range 迭代字符串时，返回每个字符的索引和 Unicode 代码点（rune）

func test2() {
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
}

func main() {
	test1()
	test2()
}
