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

	// fmt.Print  普通打印，不接受格式化操作
	fmt.Println(name, age)

	// fmt.Println 打印一行，并换行
	fmt.Println("hello world")
	fmt.Println(name, age)

	// fmt.Printf 格式化打印
	fmt.Printf("hello %s ", name)
	fmt.Printf("hello %s , age: %d ", name, age)

	// fmt.Sprintf
	info := fmt.Sprintf("hello %s , age: %d ", name, age) // 不输出结果，返回一个string 对象
	fmt.Println(info)
	fmt.Println(reflect.TypeOf(info)) //string

	// 格式化打印练习
	// 整形和浮点型
	fmt.Printf("%b\n", 12)  //1100,打印二进制
	fmt.Printf("%d\n", 12)  //12  打印不带符号的整型
	fmt.Printf("%+d\n", 12) //+12 ，打印带符号的整型
	fmt.Printf("%o\n", 12)  //14 打印不带零的八进制
	fmt.Printf("%O\n", 12)  //0o14 .打印带零的八进制
	fmt.Printf("%x\n", 12)  //c打印小写的十六进制
	fmt.Printf("%X\n", 12)  //C ， 打印大写的十六进制

	fmt.Printf("%f\n", 3.1415)   //3.141500，打印浮点型数据，默认小数点后6位，不够用0补全
	fmt.Printf("%.3f\n", 3.1415) //3.142 , 打印小数点后3位
	fmt.Printf("%e\n", 1000.25)  //1.000250e+03 ，使用科学计数法显示

	// 设置宽度 ,使用-和数字的组合，定义宽度，如下%-20s 或 %-4d
	fmt.Printf("学号：%s 姓名：%-20s 平均成绩：%-4d\n", "1001", "alvin", 100)
	fmt.Printf("学号：%s 姓名：%-20s 平均成绩：%-4d\n", "1002", "zuibangdeyuanlaoshi", 98)
	fmt.Printf("学号：%s 姓名：%-20s 平均成绩：%-4d\n", "1003", "x", 78)

}
