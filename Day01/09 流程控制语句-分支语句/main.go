package main

import "fmt"

/*
分支语句：if ,switch
*/
func main() {

	// 双分支语句
	var age = 8
	if age > 18 {
		fmt.Println("老年人")
	} else {
		fmt.Println("未成年")
	}

	// 单分支语句
	var a, b = 10, 20
	if a > b {
		a, b = b, a
	}

	//多分支语句
	if age > 10 {

	} else if age > 19 {

	}
	// switch var {
	//	 case 1 :
	//		fmt.Println("1")
	//	 case 2,3,4,5:
	//		fmt.Println("2")
	//   default:
	//      fmt.Println("非法输入")
	//}
	var choice = 3

	switch choice {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4, 5, 6:
		fmt.Println("3")
	default:
		fmt.Println("非法输入")
	}
}
