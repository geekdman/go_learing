package main

import "fmt"

/*
  返回来的函数和变量同生共死

*/

func getCounter(i int) func() {
	return func() {
		i++
		fmt.Println(i)
		fmt.Printf("%p\n",&i)
	}
}

func main() {

	counter := getCounter(100)   //其中i和 func 同生共死，一旦被创建，就会
	counter()
	counter()
	counter2 := getCounter(200)   //其中i和 func 同生共死，一旦被创建，就会
	counter2()
	counter2()

}