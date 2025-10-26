package main

import "fmt"

func test1() {
	// Go 语言数组声明需要指定元素类型及元素个数, var arrayName [size]dataType
	// 在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，对于整数类型，初始值为 0
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func test2() {
	var i, j, k int
	// 可以使用初始化列表来初始化数组的元素
	// 还可以使用 := 简短声明语法来声明和初始化数组
	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance1[%d] = %f\n", i, balance1[i])
	}

	// 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	// 如果设置了数组的长度，我们还可以通过指定下标来初始化元素
	// 将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

func main() {
	test1()
	test2()
}
