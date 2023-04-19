package main

import "fmt"

/*
字符串基本操作

	1、索引操作
	2、字符串拼接
	3、转移符
*/
func main() {
	s := "hello world"
	//1、索引操作
	//字符串[索引]，包含开始，不包含结尾
	fmt.Println(string(s[1]))
	fmt.Println(string(s[6:9])) //包含6，不包含9
	fmt.Println(string(s[6:]))  //从6位置到最后
	fmt.Println(string(s[:5]))  //从开始位置到索引5位置

	//2、字符串拼接
	var name1 = "test1"
	var name2 = "test2"
	name := name1 + name2 //使用+ 号进行字符串拼接
	fmt.Println(name)

	//3 转义符  "\" "\n"
	s1 := "yuan\n rain\n alvin"
	fmt.Println(s1)

	//4、多行字符串
	fmt.Println("床前明月光 疑是地上爽")

	s4 := `
	1、添加客户
	2、查询客户
 	3、修改客户
	4、删除客户
	5、退出程序
`
	fmt.Println(s4)
}
