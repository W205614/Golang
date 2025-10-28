package main

import (
	"fmt"
	"math"
)

// 接口（interface）是 Go 语言中的一种类型，用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约
// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
// Go 的接口设计简单却功能强大，是实现多态和解耦的重要工具
// 接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计

// 接口的特点:
// 1.隐式实现：
//   Go 中没有关键字显式声明某个类型实现了某个接口
//   只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口
// 2.接口类型变量：
//   接口变量可以存储实现该接口的任意值
//   接口变量实际上包含了两个部分：
//     动态类型：存储实际的值类型
//     动态值：存储具体的值
// 3.零值接口：
//   接口的零值是 nil。
//   一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值
// 4.空接口：
//   定义为 interface{}，可以表示任何类型

// 接口的常见用法
// 1.多态：不同类型实现同一接口，实现多态行为
// 2.解耦：通过接口定义依赖关系，降低模块之间的耦合
// 3.泛化：使用空接口 interface{} 表示任意类型

// 接口定义和实现
// 接口定义使用关键字 interface，其中包含方法声明
// 定义接口
/*type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	...
	method_namen [return_type]
}*/

// 定义结构体
/*type struct_name struct {

}*/

// 实现接口方法
/*func (struct_name_variable struct_name) method_name1() [return_type] {
	//方法实现
}*/

/*func (struct_name_variable struct_name) method_namen() [return_type] {
	// 方法实现
}*/

// 实现接口: 类型通过实现接口要求的所有方法来实现接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func test1() {
	c := Circle{Radius: 5}
	var s Shape = c // // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	test1()
}
