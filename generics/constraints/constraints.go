package main

import "fmt"

// 内置约束
// 1. any 约束
// any 是空接口 interface{} 的别名，表示任何类型都可以
func PrintAny[T any](value T) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func test1() {
	PrintAny(42)
	PrintAny("hello")
	PrintAny(3.14)
}

// 2. comparable 约束
// comparable 表示类型支持 == 和 != 操作符
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func test2() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(FindIndex(numbers, 3))

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(FindIndex(names, "Bob"))
}

// 3. 联合约束（Union Constraints）
// 使用 | 运算符组合多个类型
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func test3() {
	fmt.Println(Add(10, 20))
	fmt.Println(Add(3.14, 2.71))
}

// 自定义约束
// 1. 方法约束
// 定义需要特定方法的约束
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func test4() {
	person := Person{Name: "Alice", Age: 25}
	PrintString(person)
}

// 2. 复杂约束
// 结合类型和方法要求。
// 要求类型是数字且实现 String() 方法
type NumericStringer interface {
	Number
	String() string
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
