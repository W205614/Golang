package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

// Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数
// 同一个程序中的所有 goroutine 共享同一个地址空间
func test1() {
	go sayHello()
	for i := 0; i < 5; i++ {
		fmt.Println("Main")
		time.Sleep(100 * time.Millisecond)
	}
}

// 通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func test2() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 计算数组前半部分和
	go sum(s[len(s)/2:], c) // 计算数组后半部分和
	x, y := <-c, <-c        // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

// 通道缓冲区
// 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
// ch := make(chan int, 100)
func test3() {
	// 定义了一个缓冲区大小为2的存储整数类型的带缓冲通道
	ch := make(chan int, 2)

	// ch是带缓冲的通道，可以同时发送两个数据
	// 而不需要立刻去同步数据
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Go 遍历通道与关闭通道
func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func test4() {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个数据之后就关闭了通道
	// 所以这里我们 range 函数在接收到 10 个数据之后就结束了
	// 如果上面的 c 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了
	for i := range c {
		fmt.Println(i)
	}
}

// Select 语句
// select 语句使得一个 goroutine 可以等待多个通信操作。select 会阻塞，直到其中的某个 case 可以继续执行
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func test5() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}() // 调用匿名函数
	fibonacci2(c, quit)
}

// 使用 WaitGroup
// sync.WaitGroup 用于等待多个 Goroutine 完成
// 同步多个 Goroutine：
func worker(id int, wg *sync.WaitGroup) {
	// defer 是 Go 语言中的一个关键字，用于延迟函数的执行，直到包含 defer 的函数返回之前
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

func test6() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增减计数器
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
