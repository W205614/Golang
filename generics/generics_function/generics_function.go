package main

import "fmt"

// 1. 通用工具函数
// 交换两个值
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// 判断切片是否包含元素
func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// 去重函数
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func test1() {
	// Swap 示例
	a, b := 10, 20
	a, b = Swap(a, b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	// Contains 示例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Contains(numbers, 3))

	// Unique 示例
	duplicates := []int{1, 2, 3, 4, 4, 5}
	unique := Unique(duplicates)
	fmt.Println(unique)
}

// 2. 数学运算函数
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// 求切片最大值
func Max[T Number](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// 求切片最小值
func Min[T Number](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	min := slice[0]
	for _, value := range slice[1:] {
		if value < min {
			min = value
		}
	}
	return min
}

// 求切片平均值
func Average[T Number](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}

	var sum T
	for _, value := range slice {
		sum += value
	}

	return float64(sum) / float64(len(slice))
}

func test2() {
	ints := []int{1, 5, 3, 9, 2}
	floats := []float64{1.1, 5.5, 3.3, 9.9, 2.2}

	fmt.Printf("Max int: %d\n", Max(ints))
	fmt.Printf("Min float: %.1f\n", Min(floats))
	fmt.Printf("Average: %.2f\n", Average(floats))
}

func main() {
	test1()
	test2()
}
