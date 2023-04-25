package main

import "fmt"

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
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	//
	//for i :=0 ;i<len(a);i++ {
	//	v :=
	//}

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v }
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
