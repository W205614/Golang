package main

import "fmt"

// 声明变量的一般形式是使用 var 关键字
// 可以一次声明多个变量
func test1() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
}

// 变量声明
// 第一种，指定变量类型，如果没有初始化，则变量默认为零值
func test2() {

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
}

func test3() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// 第二种，根据值自行判定变量类型
func test4() {
	var d = true
	fmt.Println(d)
}

// 第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
func test5() {
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)
}

// 多变量声明
// 如果你声明了一个局部变量却没有在相同的代码块中使用它，会得到编译错误
// 但是全局变量是允许声明但不使用的
// 声明全局变量
// 1. 使用默认值
var x, y int

// 2. 使用因式分解关键字的写法
var (
	a int
	b bool
)

// 3. 同时初始化多个变量
var c, d int = 1, 2

// 4. 省略类型并初始化多个变量
var e, f = 123, "hello"

// intVal := 1 相等于：
// var intVal int
// intVal = 1

func test6() {
	g, h := 123, "hello"
	_, i := 4, 7 // 空白标识符 _ 也被用于抛弃值，如值 4 在：_, i := 4, 7 中被抛弃
	fmt.Println(x, y, a, b, c, d, f, g, h, i)
	fmt.Println(&x)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
