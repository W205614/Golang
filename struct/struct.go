package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func test_books_struct1() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	// 使用 key->value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	// 忽略的字段使用默认值
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

func test_books_struct2() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495700

	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	// fmt.Printf("Book1 title: %s\n", Book1.title)
	// fmt.Printf("Book1 author: %s\n", Book1.author)
	// fmt.Printf("Book1 subject: %s\n", Book1.subject)
	// fmt.Printf("Book1 bppk_id: %d\n", Book1.book_id)

	// fmt.Printf("Book2 title: %s\n", Book2.title)
	// fmt.Printf("Book2 author: %s\n", Book2.author)
	// fmt.Printf("Book2 subject: %s\n", Book2.subject)
	// fmt.Printf("Book2 bppk_id: %d\n", Book2.book_id)

	printBook1(Book1)
	printBook1(Book2)

	printBook2(&Book1)
	printBook2(&Book2)
}

// 结构体作为函数参数
func printBook1(book Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book bppk_id: %d\n", book.book_id)
}

// 结构体指针
func printBook2(book *Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book bppk_id: %d\n", book.book_id)
}

func main() {
	test_books_struct1()
	test_books_struct2()
}
