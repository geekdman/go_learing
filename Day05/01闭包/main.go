package main

import "fmt"

/*
  返回来的函数和变量同生共死

*/

func getCounter() func() {
	var i =0
	return func() {
		i++
		fmt.Println("%p\n",&i)
	}
}

func main() {

	counter := getCounter()   //其中i和 func 同生共死，一旦被创建，就会
	counter()   //1
	counter()   //2
	counter1 := getCounter()
	counter1()  //1
	counter1()  //2
}