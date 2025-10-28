package main

import (
	"fmt"
	"strconv"
)

// 数值类型转换
func test1() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

// 字符串类型转换
// 将字符串转换为整数
func test2() {
	str := "123"
	// strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转化为整数为: %d\n", str, num)
	}
}

// 将整数转换为字符串
func test3() {
	num := 123
	str := strconv.Itoa(num)
	fmt.Printf("整数 %d 转化为字符串为: '%s'\n", num, str)
}

// 将字符串转换为浮点数
func test4() {
	str := "3.14"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转化错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转化为浮点型为: %f\n", str, num)
	}
}

// 将浮点数转换为字符串
func test5() {
	num := 3.14
	str := strconv.FormatFloat(num, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为: '%s'\n", num, str)
}

// 接口类型转换有两种情况：类型断言和类型转换
// 1.类型断言
// 类型断言用于将接口类型转换为指定类型
// value.(type) 或者 value.(T), 其中 value 是接口类型的变量, type 或 T 是要转换成的类型
// 如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功
func test6() {
	// 接口类型变量 i
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

// 2.类型转换
// 类型转换用于将一个接口类型的值转换为另一个接口类型
// T(value), T 是目标接口类型，value 是要转换的值
// 在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错
// 定义一个接口 Writer
type Writer interface {
	Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
	str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func test7() {
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}
	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)
	// 修改 StringWriter 的字段
	sw.str = "Hello World"
	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)
}

// 空接口 interface{} 可以持有任何类型的值
func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func test8() {
	printValue(42)
	printValue("hello")
	printValue(3.14)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}
