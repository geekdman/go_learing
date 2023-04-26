package main

import "fmt"

/*
解析执行过程
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] // [10, 20]
	s2 := s1       //  // [10, 20]
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(arr)
*/

}
func main() {
	arr := [4]int{10, 20, 30, 40}       //初始化一个长度为4 的数组arr，
	s1 := arr[0:2] // [10, 20]          //对数组arr 执行切片操作，得到切片s1，此时s1 是一个长度为2，容量为4的切片，内容为[10,20]
	s2 := s1       //  // [10, 20]      //从切片s1 创建s2 。s2 和s1 的底层数组为同一个。s2 也是一个长度为2 ，容量为4的切片，内容为[10,20]

	/*	执行第一次apend(s1,1) ,此时s1 的长度为3，容量为4，内容为[10,20,1];因为s2 和s1 是同一个底层数组，所以此时s2 和s1 一样
		执行第二次apend(apend(s1,1),2),由于第一次apend 得到的切片，

	*/
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(arr)
}