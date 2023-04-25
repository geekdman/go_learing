package main

import "fmt"

/* 作业4
	切片的反转 s:= []int{1,2,3,5,4}
*/
func main() {
	s:= []int{1,2,3,5,4}
	var tmp = make([]int,len(s))
	for i,v := range s {
		tmp[len(s) -i -1]= v
	}
	s = tmp
	fmt.Println(s)
}