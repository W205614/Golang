package main

import (
	"fmt"
	"time"
)

func test_for1() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 当 i 等于 5 时跳出循环
		}
		fmt.Println(i)
	}
}

func test_for2() {
	var a int = 10

	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			break
		}
	}
}

// re 是一个标签，用于标识外层的 for 循环, 通过在 break 语句后面加上标签 re，可以使 break 跳出外层的循环，而不仅仅是内层循环
func test_for3() {
	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}

func test_switch() {
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
		break // 跳出 switch 语句
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	}
}

func test_select1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("Received from ch1.")
	case <-ch2:
		fmt.Println("Received from ch2.")
		break // 跳出 select 语句
	}
}

// select 语句中使用 return 来提前结束执行的情况
func process(ch chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println("Received value:", val)
			// 执行一些逻辑
			if val == 5 {
				return // 提前结束 select 语句的执行
			}
		default:
			fmt.Println("No value received yet.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func test_select2() {
	ch := make(chan int)

	go process(ch)

	time.Sleep(2 * time.Second)
	ch <- 1
	time.Sleep(1 * time.Second)
	ch <- 3
	time.Sleep(1 * time.Second)
	ch <- 5
	time.Sleep(1 * time.Second)
	ch <- 7

	time.Sleep(2 * time.Second)
}

func main() {
	test_for1()
	test_for2()
	test_for3()
	test_switch()
	test_select1()
	test_select2()
}
