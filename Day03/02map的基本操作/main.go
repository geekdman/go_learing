package main

import "fmt"

func main() {

	var stuSlice = []string{"tom", "32", "male"}
	var stuMap = map[string]string{"name": "tony", "age": "32"}

	// 1、查看键值
	fmt.Println(stuSlice[1])
	fmt.Println(stuMap["name"])

	for key, val := range stuMap {
		fmt.Println(key, val)
	}

	// 2、删除键值
	delete(stuMap, "age")
	fmt.Println(stuMap)

	// 3、添加键值对和更新,key 存在，就是修改操作； key不存在，就是添加操作
	stuMap["name"] = "李四"
	stuMap["weight"] = "80kg"
	fmt.Println(stuMap)

}
