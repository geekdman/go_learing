package main

import "fmt"

/*
1、01指针类型
指针的基本使用
计算机所有的数据都必须存放在内存中，数据在内存中的地址为指针

提取变量的地址
var x = 100
var p =&x   p的数据类型是 *int, 是根据p指向的x的值的类型，来确定的
*p

2、 new 函数  开辟空间
================
指针默认为nil，不能直接使用，在定义指针的时候，所以最好使用new 函数，为指针初始化一块地址
*/
func main() {
	var a int
	var p *int // 定义指针变量
	p = &a
	v := *p
	//	获取x 的地址
	fmt.Println(a) //0
	fmt.Println(v)
	fmt.Println(&a) //0xc0000a6058
	fmt.Println(p)  //0xc0000a6058

	//练习题
	var x = 10
	var y = &x
	var z = &y
	fmt.Println(x)
	fmt.Println(*y)
	fmt.Println(z)

	//练习题
	var x1 = 100
	var y1 = &x1 // y1 的类型是 *int
	var z1 = &y1 //z1 的类型是**int
	**z1 = 200
	fmt.Println(x1)

	// new 函数和make函数
	//指针默认为nil，不能直接使用，在定义指针的时候，所以最好使用new 函数，为指针初始化一块地址
	var p1 *int //默认值为nil ，没有开辟内存地址空间，需要使用new() 函数创建内存地址空间
	fmt.Println(p)
	p1 = new(int)
	*p1 = 100
}
