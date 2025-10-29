package main

import (
	"fmt"
	"sync"
)

// 1. 泛型结构体
// 泛型栈实现
type Stack[T any] struct {
	elements []T
}

// 入栈
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// 出栈
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]

	return value, true
}

// 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	return s.elements[len(s.elements)-1], true
}

// 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func test1() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop())
}

// 2. 泛型映射（Map）
// 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// 设置键值对
func (m *SafeMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = value
}

// 获取值
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	value, exists := m.data[key]
	return value, exists
}

// 删除键
func (m *SafeMap[K, V]) Delete(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}

// 获取所有键
func (m *SafeMap[K, V]) Keys() []K {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

func test2() {
	scores := NewSafeMap[string, int]()
	scores.Set("Alice", 95)
	scores.Set("Bob", 87)

	if score, exists := scores.Get("Alice"); exists {
		fmt.Printf("Alice's score: %d\n", score)
	}
	fmt.Println("Keys:", scores.Keys())
}

func main() {
	test1()
	test2()
}
