package main

import "fmt"

// Map 是无序的，遍历 Map 时返回的键值对的顺序是不确定的
// 在获取 Map 的值时，如果键不存在，返回该类型的零值
// Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量
func test1() {
	// 创建集合
	var siteMap map[string]string
	siteMap = make(map[string]string)

	// key - value
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	// 使用键输出value
	for k := range siteMap {
		fmt.Println(k, "首都是", siteMap[k])
	}

	// 查看元素在集合中是否存在
	k, v := siteMap["Facebook"]
	if v {
		fmt.Println("Facebook 的站点是", k)
	} else {
		fmt.Println("Facebook 站点不存在")
	}
}

// delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
func test2() {
	counttyCapitalMap := map[string]string{"France": "Pairs", "Italy": "Rome", "Japen": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	for country := range counttyCapitalMap {
		fmt.Println(country, "首都是", counttyCapitalMap[country])
	}

	delete(counttyCapitalMap, "France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除元素后地图")

	for country := range counttyCapitalMap {
		fmt.Println(country, "首都是", counttyCapitalMap[country])
	}
}

func main() {
	test1()
	test2()
}
