package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// os 包含了很多系统操作包

func ReadBytes(file *os.File) {
	//按照byte读取
	data := make([]byte, 6)
	file.Read(data)
	fmt.Println(data)
	//byte 转换成 string
	fmt.Println(string(data))
}

// 按照line读
func ReadByLine(file *os.File) {
	// 按照换行符号

	reader := bufio.NewReader(file)
	/*	data, _, _ := reader.ReadLine()
		fmt.Println(string(data))*/
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

func main() {

	file, err := os.Open("满江红.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 按照行读取文件
	ReadByLine(file)
	// 按照字节读取文件
	ReadBytes(file)

	// 直接全部读取到内存里，适用于小文件
	fmt.Println(os.ReadFile("满江红.txt"))

}
