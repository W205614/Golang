package main

import "fmt"

// 1. 组合（Composition）
// 组合是 Go 中实现代码复用的主要方式。通过将一个结构体嵌入到另一个结构体中，子结构体可以"继承"父结构体的字段和方法
// 父结构体
type Animal1 struct {
	Name string
}

// 父结构体的方法
func (a *Animal1) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog1 struct {
	Animal // 嵌入 Animal 结构体
	Breed  string
}

func test1() {
	dog := Dog1{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	dog.Speak() // 调用父结构体的方法
	fmt.Println("Breed:", dog.Breed)
}

// 2. 接口（Interface）
// 接口是 Go 中实现多态的主要方式。通过定义接口，不同的结构体可以实现相同的方法，从而实现类似继承的多态行为
// 定义接口
type Speaker interface {
	Speak()
}

// 父结构体
type Animal struct {
	Name string
}

// 实现接口方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
	Animal
	Breed string
}

func test2() {
	var speaker Speaker

	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	speaker = &dog
	speaker.Speak() // 通过接口调用方法
}

// 完整继承模拟：
// 基类
type Vehicle struct {
	Brand string
}

func (v *Vehicle) Start() {
	fmt.Println(v.Brand, "started")
}

// 派生类
type Car struct {
	Vehicle // 嵌入Vehicle
	Model   string
}

// 重写Start方法
func (c *Car) Start() {
	fmt.Println(c.Brand, c.Model, "car started")
}

func test3() {
	v := Vehicle{Brand: "Toyota"}
	c := Car{
		Vehicle: Vehicle{Brand: "Honda"},
		Model:   "Civic",
	}

	v.Start()         // Toyota started
	c.Start()         // Honda Civic car started
	c.Vehicle.Start() // Honda started
}

func main() {
	test1()
	test2()
	test3()
}
