package main

import "fmt"

/*
go 语言只有for

	  格式: for condistion {

	}
	使用break 退出循环
	使用continue ：退出当次循环
*/
func main() {
	//for true {
	//	fmt.Println("hello world")
	//}

	sum := 0
	for n := 1; n <= 100; n++ {
		sum += n
	}
	fmt.Println(sum)
}
