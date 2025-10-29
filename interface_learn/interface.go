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

// 空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集
//
//	任意类型都实现了空接口。
//	常用于需要存储任意类型数据的场景，如泛型容器、通用参数等
func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func test2() {
	printValue(42)
	printValue("hello")
	printValue(3.14)
	printValue([]int{1, 2})
}

// 类型断言用于从接口类型中提取其底层值
// value := iface.(Type), iface 是接口变量, Type 是要断言的具体类型, 如果类型不匹配，会触发 panic
func test3() {
	var i interface{} = "hello"
	str := i.(string) // 类型断言
	fmt.Println(str)
}

// 为了避免 panic，可以使用带检查的类型断言：
// value, ok := iface.(Type), ok 是一个布尔值，表示断言是否成功, 如果断言失败，value 为零值，ok 为 false
func test4() {
	var i interface{} = 42
	if str, ok := i.(string); ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("Not a string")
	}
}

// 类型选择（type switch）
// type switch 是 Go 中的语法结构，用于根据接口变量的具体类型执行不同的逻辑
func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func test5() {
	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})
}

// 接口组合
// 接口可以通过嵌套组合，实现更复杂的行为描述
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "Reading data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}

func test6() {
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello Go!")
}

// 动态值和动态类型
func test7() {
	var i interface{} = 42
	fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}

// 接口的零值是 nil
// 当接口变量的动态类型和动态值都为 nil 时，接口变量为 nil
func test8() {
	var i interface{}
	fmt.Println(i == nil)
}

// 接口的使用练习
// 练习1
type Phone interface {
	call()
}

type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct{}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func exercise1() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

// 练习2
type Shape1 interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle1 struct {
	radius float64
}

func (c Circle1) area() float64 {
	return 3.14 * c.radius * c.radius
}

func exercise2() {
	var s Shape1

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Circle1{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
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
	exercise1()
	exercise2()
}
