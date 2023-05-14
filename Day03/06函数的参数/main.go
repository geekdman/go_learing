package main

import "fmt"

/*
带参数的函数

func 函数名字(parameter type) {}

*/

func PrintLoop(end int) {

	for i := 1; i < end; i++ {
		fmt.Println(i)
	}

}
func main() {
	PrintLoop(3)
	PrintLoop(5)
}
