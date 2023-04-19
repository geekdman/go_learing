package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"

	// 小写转大写 ToUpper(s)
	sUpper := strings.ToUpper(s)
	fmt.Println(s)
	fmt.Println(sUpper)

	// 大写转小写 ToLower(s)
	sLower := strings.ToLower(s)
	fmt.Println(sLower)

	// 判断以什么开头
	strings.HasPrefix(s, "张三") //false

	//判断以什么结尾
	strings.HasSuffix(s, "rld") //true

	//判断包含什么
	strings.Contains(s, "or")

	//去掉两端的内容
	username := "  tony"
	fmt.Println(len(username))
	fmt.Println(username)
	username_1 := strings.Trim(username, " ")
	fmt.Println(len(username_1))
	fmt.Println(username_1)

	// 索引函数,从左到右，第一个匹配的
	s1 := "hello test world test "
	fmt.Println(strings.Index(s1, "test"))

	// 分割字符串
	citys := "北京 上海 南京"
	countrys := "五台山 太原 大同"

	strings.Split()
	// 拼接字符串
	s6 := []string{"foo", "bar", "baz"}
	strings.Join(s6, ",")
}
