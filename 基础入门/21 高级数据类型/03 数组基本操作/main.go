package main

import "fmt"
func main() { // 1索引
//	取值
//	循环
//	遍历
	var names [4]string
	var names= [4]string["张三","李四","王五","赵六"]
	for i := range names {
		fmt.Println(names)
	}
}

