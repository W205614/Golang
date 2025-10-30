package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 文件创建
// 在 Go 语言中，我们使用 os 包来创建文件。
// os.Create 函数用于创建一个文件，并返回一个 *os.File 类型的文件对象。创建文件后，我们通常需要调用 Close 方法来关闭文件，以释放系统资源
func test1() {
	// 创建文件，如果文件已存在会被截断（清空）
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("文件创建成功")
}

// 文件的打开与关闭
// 在 Go 语言中，我们使用 os 包来打开和关闭文件。
// os.Open 函数用于打开一个文件，并返回一个 *os.File 类型的文件对象。打开文件后，我们通常需要调用 Close 方法来关闭文件，以释放系统资源
func test2() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opend successfully!")
}

// 文件的读取
// Go 语言提供了多种读取文件的方式，包括逐行读取、一次性读取整个文件等。我们可以使用 bufio 包来逐行读取文件，或者使用 ioutil 包来一次性读取整个文件
// 逐行读取文件
func test3() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 使用 bufio.NewScanner 创建了扫描器
	scanner := bufio.NewScanner(file)
	// 通过 scanner.Scan() 逐行读取文件内容
	for scanner.Scan() {
		// 使用 scanner.Text() 获取每一行的文本
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// 一次性读取整个文件
func test4() {
	// 使用 os.ReadFile 一次性读取整个文件的内容，并将其转换为字符串打印出来
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}

// 文件的写入
// Go 语言也提供了多种写入文件的方式，包括逐行写入、一次性写入等。我们可以使用 os 包来创建和写入文件
func test5() {
	file, err := os.Create("Write1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("直接写入字符串\n")

	data := []byte("写入字节切片\n")
	file.Write(data)

	fmt.Fprintf(file, "格式化写入: %d\n", 123)
}

// 逐行写入文件
func test6() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 使用 bufio.NewWriter 创建一个写入器
	writer := bufio.NewWriter(file)
	// 使用 fmt.Fprintln 将字符串写入文件，并调用 writer.Flush() 确保所有数据都被写入文件
	fmt.Fprintln(writer, "Hello, World!")
	writer.Flush()
}

// 一次性写入文件
func test7() {
	content := []byte("Hello, World!")
	err := os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully!")
}

// 文件的追加写入
// 有时候我们需要在文件的末尾追加内容，而不是覆盖原有内容。Go 语言提供了 os.OpenFile 函数来实现这一功能
func test8() {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("Append text\n"); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Text appended successfully")
}

// 文件的删除
// 在 Go 语言中，我们可以使用 os.Remove 函数来删除文件
func test9() {
	err := os.Remove("output.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully!")
}

// 文件信息与操作
// 获取文件信息
func test10() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size(), "字节")
	fmt.Println("权限:", fileInfo.Mode())
	fmt.Println("最后修改时间:", fileInfo.ModTime())
	fmt.Println("是目录吗:", fileInfo.IsDir())
}

// 检查文件是否存在
func test11() {
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}
}

// 重命名和移动文件
func test12() {
	err := os.Rename("test.txt", "test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("重命名成功")
}

// 目录操作
// 创建目录
func test13() {
	// 创建单个目录
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 递归创建多级目录
	err = os.MkdirAll("path/to/newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// 读取目录内容
func test14() {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(),
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"))
	}
}

// 删除目录
func test15() {
	// 删除空目录
	err := os.Remove("newdir")
	if err != nil {
		log.Fatal(err)
	}

	// 递归删除目录及其内容
	err = os.RemoveAll("path/to/newdir")
	if err != nil {
		log.Fatal(err)
	}
}

// 高级文件操作
// 文件复制
func test16() {
	srcFile, err := os.Open("Write1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("destination.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("复制完成，共复制 %d 字节", bytesCopied)
}

// 文件追加
func test17() {
	file, err := os.OpenFile("Write1.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString("新的日志内容\n"); err != nil {
		log.Fatal(err)
	}
}

// 临时文件和目录
func test18() {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // 清理

	fmt.Println("临时文件:", tmpFile.Name())

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // 清理

	fmt.Println("临时目录:", tmpDir)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	// test14()
	// test15()
	// test16()
	// test17()
	test18()
}
