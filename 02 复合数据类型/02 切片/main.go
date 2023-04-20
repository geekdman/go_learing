package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
	切片 slice 定义：
	slice 特点：
		1、长度可变
	*/
	// 方式一： var s1 []type , 其中type 为基本数据类型
	var s1 []int
	fmt.Println(reflect.TypeOf(s1))   //[]int

	// 比较： 定义一个数组
	//var s2 [4]int
	//fmt.Println(reflect.TypeOf(s2))   //[4]int
}