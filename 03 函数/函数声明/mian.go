package main

import (
	"fmt"
	"os"
)

/*
 函数声明，包含函数名称、一个形参列表、一个可选的返回列表以及函数体
格式：
func name(parameter-list) (result-list){ }
*/

func added(a int, b int) int {
	return a + b
}
func divided(a int, b int) int {
	return a / b
}

func main() {
	a := 1
	b := 2
	fmt.Println(added(a, b))
	fmt.Println(fmt.Scan(os.Args))
}
