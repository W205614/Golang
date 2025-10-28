package main

import (
	"fmt"           // 用于格式化输入和输出
	"os"            // 提供了与操作系统交互的功能
	"path/filepath" // 提供了跨平台的路径操作功能
)

// 递归计算阶乘
func test1(n int) int {
	// 基准条件
	if n == 0 {
		return 1
	}

	// 递归条件
	return n * test1(n-1)
}

// 斐波那契数列
func test2(n int) int {
	if n < 2 {
		return n
	}

	return test2(n-2) + test2(n-1)
}

// 求平方根
// 计算平方根的数, 当前的猜测值, 上一次的猜测值, 精度阈值（用于判断是否达到足够的精度）
func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	// 如果 diff 的绝对值小于精度阈值 epsilon，则认为当前猜测值已经足够接近真实的平方根，返回当前猜测值 guess
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		return guess
	}

	// 牛顿迭代法的标准公式，用于更快地逼近平方根
	newGuess := (guess + x/guess) / 2
	// 如果新的猜测值与上一次的猜测值相同，说明已经收敛，返回当前猜测值 guess
	if newGuess == prevGuess {
		return guess
	}

	return sqrtRecursive(x, newGuess, guess, epsilon)
}

func sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}

// 文件目录遍历
// 当前要遍历的目录路径, 用于缩进的字符串，表示当前目录的层级
func walkDir(dir string, indent string) {
	// 使用 os.ReadDir 读取目录中的所有条目（文件和子目录）
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Failed to read directory %s: %v\n", dir, err)
		return
	}

	for _, entry := range entries {
		// 打印当前条目的名称，并根据 indent 参数进行缩进
		fmt.Println(indent + entry.Name())
		// 如果当前条目是一个目录（entry.IsDir() 返回 true），则递归调用 walkDir 函数，将当前目录路径与子目录名称拼接，并增加缩进
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), indent+"")
		}
	}
}

func main() {
	fmt.Println(test1(5))

	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", test2(i))
	}
	fmt.Println()

	x := 25.0
	result := sqrt(x)
	fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)

	walkDir(".", "")
}
