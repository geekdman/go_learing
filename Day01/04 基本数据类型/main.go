package main

import (
	"fmt"
	"reflect"
)

/*
基本数据类型：
	1、整型
	2、浮点型
	3、布尔类型
	4、字符串类型
*/

func main() {
	// 浮点类型
	var f1 = 3.14
	fmt.Println(reflect.TypeOf(f1))
	var f2 = 6.2e-2
	fmt.Println(reflect.TypeOf(f2))

	// 声明布尔类型
	var b1 = true
	var b2 = false
	fmt.Println(reflect.TypeOf(b1))
	fmt.Println(reflect.TypeOf(b2))

	var x, y = 100, 200
	var b3 = y > x
	fmt.Println(b3, reflect.TypeOf(b3))

	// 字符串
	var s = "test" //go 语言只能用双引号
	var c1 = 'a'   //单引号只能包裹一个字符
}
