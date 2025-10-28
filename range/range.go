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

// 映射（Map）
// for 循环的 range 格式可以省略 key 和 value
func test3() {
	// 创建一个空的 map，key 是 int 类型，value 是 float32 类型
	map1 := make(map[int]float32)

	// 向 map1 中添加 key-value 对
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 遍历 map1，读取 key 和 value
	for key, value := range map1 {
		// 打印 key 和 value
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 遍历 map1，只读取 key
	for key := range map1 {
		// 打印 key
		fmt.Printf("key is: %d\n", key)
	}

	// 遍历 map1，只读取 value
	for _, value := range map1 {
		// 打印 value
		fmt.Printf("value is: %f\n", value)
	}
}

// 通道（Channel）
// range 遍历从通道接收的值，直到通道关闭
func test4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// 忽略值
// 在遍历时可以使用 _ 来忽略索引或值
func test5() {
	nums := []int{2, 3, 4}

	// 忽略索引
	for _, num := range nums {
		fmt.Println("value:", num)
	}

	// 忽略值
	for i := range nums {
		fmt.Println("index:", i)
	}
}

// range 遍历其他数据结构
func test6() {
	//这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range 也可以用在 map 的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
