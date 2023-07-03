package main

import (
	"fmt"
	"reflect"
)

/*
结构体
1、定义：
type 类型名 struct {
    字段1  字段1类型
	字段2  字段2类型
    ...
}

2、实例化
方式1：
 var p1 Person
*/

type  Person struct {
	// name , agem, isMarried,socres 是成员变量
	name string
	age int
	isMarried bool
	socres map[string]int
}

func main()  {
	// 实例化1
	var p1 Person
	p1.name="zhangsan"
	p1.age = 18
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.isMarried)
	fmt.Println(p1.socres)

	// 结构体的存储是连续的,成员变量的地址是连续开辟的
	fmt.Println(&p1)    //p1 在内存的中的地址，和第一个成员变量的地址一样
	fmt.Println(&p1.name)
	fmt.Println(&p1.age)
	fmt.Println(&p1.isMarried)
	fmt.Println(&p1.socres)

	// 实例化 2,按照顺序赋值
	var  p2 = Person{"lisi",18,true, map[string]int{}}
	fmt.Println(p2)
	// 实例化3 ，按照键进行赋值，可以选着初始化哪些变量。
	var p3 =  Person{name: "lisi",socres: map[string]int{}}
	fmt.Println(p3)

	// 实例化3 ，实例化取址
	var p4 = &Person{name: "王五",age: 18}
	fmt.Println(p4)
	fmt.Println(reflect.TypeOf(p4))
	fmt.Println((*p4).name)  //王五
	fmt.Println(p4.name)   // 王五，在go语言中当为结构体地址的时候，访问结构其的属性，可以省略* 号，即 p.name 等于(*p).name

}
