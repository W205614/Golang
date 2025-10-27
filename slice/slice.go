package main

import "fmt"

// len() 和 cap() 函数
func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)

}

func test1() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}

// 一个切片在未初始化之前默认为 nil，长度为 0
func test2() {
	var numbers []int
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

// 切片截取
func test3() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)
	// 打印子切片[索引1，索引4)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	// 默认下限为 0
	fmt.Println("numbers[:3] ==", numbers[:3])
	// 默认上限为 len
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

// append() 和 copy() 函数
func test4() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	// 当append(list, [params])，先判断 list 的 cap 长度是否大于等于 len(list) + len([params])
	// 如果大于那么 cap 不变，否则 cap = 2 * max{cap(list), cap[params]}
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	// 拷贝 numbers 的内容到 numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
