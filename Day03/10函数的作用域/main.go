package main

import "fmt"

func foo() {
	x = 1000
	fmt.Println(x) //1000
}

var x = 100 // 全局变脸，只能这么定义，不能简写 x:=100

func main() {
	foo()          //1000
	fmt.Println(x) //1000
	x := 200
	fmt.Println(x) //200
	foo()          //1000
	fmt.Println(x) //200
}
