// iota，特殊常量，可以认为是一个可以被编译器修改的常量
// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
package main

import "fmt"

// iota 在 const 块中用于生成连续的整数值
// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
// 显式赋值不会将 iota 的值赋给该常量，但 iota 的计数不会重置，而是继续自增
// 如果某个常量没有显式赋值，它会继承前一个常量的值
func test1() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

// iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思）
// i << n = i * (2 ^ n)
const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func test2() {
	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)
	fmt.Println("l =", l)
}

func main() {
	test1()
	test2()
}
