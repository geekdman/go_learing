package main

import (
	"fmt"
	"time"
)

/*
  返回来的函数和变量同生共死

*/
func bar()  {
	fmt.Println("bar")
}


func getTimer(f func()) func() {
	return func() {
		start:=time.Now().Unix()
		f()
		end := time.Now().Unix()
		fmt.Println(end-start)
	}
}

func main() {
	bar := getTimer(bar)
	bar()
}