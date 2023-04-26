package main

import (
	"fmt"
)

/*
作业5 ：解析过程
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
	}
	r[i] = v }
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
*/

func main() {
	var a = [5]int{1, 2, 3, 4, 5} //声明长度为5 ，内容为1，2，3，4，5 的数组a
	var r [5]int                  //声明 长度为5 的数组r
	fmt.Printf("%v\n",&a)
	//for i, v := range &a {
	for i, v := range &a { // 循环数组a 的地址
		fmt.Printf("%p\n",&a)
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		//fmt.Println(&v)
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
