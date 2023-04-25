package main

import "fmt"

/*1. 解析执行过程(描述过程)
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
*/

func main() {
	s := make([]int, 5)             // 创建一个slice，并在内存中开辟一块长度为5、容量为5的int 类型的slice。
	s = append(s, 1, 2, 3)  // 在内存中创建一个长度为5，容量为10的int 类型的slice，并将s 的内容复制过来。得到一个新的int 类型的slice。 然后将新的slice 赋值给 变量s。
	fmt.Println(s)
}