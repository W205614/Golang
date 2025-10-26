package main

import "fmt"

// Printf 根据格式化参数生成格式化的字符串并写入标准输出
func main() {
	var stockcode = 123
	var enddate = "2025-10-23"
	var url = "Code = %d &endDate = %s"
	fmt.Printf(url, stockcode, enddate)
}
