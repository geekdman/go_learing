package main

import (
	"fmt"
	"reflect"
)

/*
定义：
	数组是具有固定长度，并且 拥有零个或者多个相同数据类型元素的序列。由于数组的长度固定，所以在Go里面很少直接使用。
特点：
	1、固定长度，长度不可变
	2、相同数据类型
*/

func main() {
	////方式一：只声明数组
	//var a1 [3]int //初始化a1是长度为3的数组。
	//var a []int
	//var b []string
	//var c []int8
	////方式二: 声明数组长度为3，并初始化数组内容为1,2,3。
	//array1 := [3]int{1, 2, 3}
	////方式三: 如果省略号 "..."出现在数组长度的位置，那么数组长度由初始化数组的元素个数决定
	//array2 := [...]int{1, 2, 3}
	////方式四：
	//array3 := [...]int{99: -1}
	s1 := []int{1,2,3,4}
	fmt.Println(len(s1),cap(s1))
	s1=append(s1,1000)
	fmt.Println(len(s1),cap(s1))

	var s3 [3][]int
	fmt.Println(s3)
	s3[0] =s1
	fmt.Println(reflect.TypeOf(s3))
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(s3[0][2]))
	fmt.Println(len(s3),cap(s3))
	fmt.Println(len(s1),cap(s1))
	//fmt.Println(len(array3))
	//
	//fmt.Println(s)
}
