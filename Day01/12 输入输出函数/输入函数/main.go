package main

import (
	"fmt"
	"reflect"
)

/*
	输出函数：主要是用来将go程序的结果输出到控制台的函数，通常是语言本身内置函数,在go语言中，定义在fmt 包中，需要先导入fmt包，才能使用
	go 语言常见的几种函数：
	fmt.Print   //输出到控制台,不接受任何格式化操作
    fmt.Println //输出到控制台并换行
    fmt.Printf  //只可以打印出格式化的字符串，只可以直接输出字符串类型的变量（不可以输出别的类型）
    fmt.Sprintf //格式化并返回一个字符串而不带任何输出
*/

func main() {
	var name string = "tony"
	var age int = 16

	// fmt.Print
	fmt.Println(name, age)

	// fmt.Println
	fmt.Println("hello world")
	fmt.Println(name, age)

	// fmt.Printf
	fmt.Printf("hello %s ", name)
	fmt.Printf("hello %s , age: %d ", name, age)

	// fmt.Sprintf
	info := fmt.Sprintf("hello %s , age: %d ", name, age) // 不输出结果，返回一个string 对象
	fmt.Println(info)
	fmt.Println(reflect.TypeOf(info)) //string

}
