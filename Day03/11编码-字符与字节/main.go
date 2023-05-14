package main

import "fmt"

/*
字符类型声明和定义
格式：
	var b1 byte
	b1 = 'a'

 备注： 其中单引号中，包裹的只能是单个字符，单个字符可以是单个英文字符，或者是单个中文字符
其中英文字符，类型是byte 或者是uint8， 中文字符是rune 或者是int32
*/

func main() {

	// 定义 byet 类型 ， 其中 byte = uint8
	/*	var s byte
		s = 'a'
		fmt.Println(reflect.TypeOf(s)) //byte
		fmt.Println(s)                 //97
		fmt.Println(string(s))         //a
	*/
	// 定义uint8
	/*	var b2 uint8
		b2 = 65
		fmt.Println(b2)         //65
		fmt.Println(string(b2)) //A
	*/

	// 定义int32  ，其中rune = int32
	var b3 rune //rune 用来存放一个单独的中文
	b3 = '中'
	fmt.Println(b3)         // 20013
	fmt.Println(string(b3)) //中
}
