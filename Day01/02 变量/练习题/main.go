package main

import "fmt"

/*
交换两个变量的值
*/
func main() {
	// 方式一
	var v1, v2 = 10, 20
	tmp := v1
	v1 = v2
	v2 = tmp
	fmt.Println(v1, v2)

	//方式二
	var v3, v4 = 11, 22
	v3, v4 = v4, v3
	fmt.Println(v3, v4)
}
