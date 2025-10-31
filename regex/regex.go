package main

import (
	"fmt"
	"regexp"
)

// 检查字符串是否匹配正则表达式
func test1() {
	pattern := `^[a-zA-Z0-9]+$`
	regex := regexp.MustCompile(pattern)

	str := "Hello123"
	if regex.MatchString(str) {
		fmt.Println("字符串匹配正则表达式")
	} else {
		fmt.Println("字符串不匹配正则表达式")
	}
}

// 查找匹配的字符串
func test2() {
	pattern := `\d+`
	regex := regexp.MustCompile(pattern)

	str := "我有 3 个苹果和 5 个香蕉"
	matches := regex.FindAllString(str, -1)
	fmt.Println("找到的数字：", matches)
}

// 替换匹配的字符串
func test3() {
	pattern := `\s+`
	regex := regexp.MustCompile(pattern)

	str := "Hello    World"
	result := regex.ReplaceAllString(str, " ")
	fmt.Println("替换后的字符串：", result)
}

// 分割字符串
func test4() {
	pattern := `,`
	regex := regexp.MustCompile(pattern)

	str := "apple,banana,orange"
	parts := regex.Split(str, -1)
	fmt.Println("分割后的字符串：", parts)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
