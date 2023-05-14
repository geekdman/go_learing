package main

import "fmt"

/*
map 类型
map 是一种通过key来获取value的一个数据结构，其底层存储方式为数组，再存储时key 不重复，当key重复时，value进行覆盖，我们通过key进行
1、如何声明?
```go
var map_name map[key_type]value_type
// `map_name` 为 map 的变量名。
// `key_type`为键类型。
// `value_type`是键对应的值类型。
```
2、

*/

func main() {
	// []interface{} , 其中interface 是关键字，可以接受任意类型
	//切片实现
	str := []interface{}{"user01", 22, "180cm"}
	fmt.Println(str)

	//map 实现
	// 方式1： 先声明再初始化
	var user02 = make(map[string]string)
	user02["name"] = "user01"
	user02["age"] = "32"
	user02["gender"] = "male"

	fmt.Println(user02)

	//方式二： 声明并初始化
	var str03 = map[string]string{"name": "user03", "age": "32", "gender": "male"}
	fmt.Println(str03)

	// 案例

	var str04 = make(map[interface{}]interface{})
	str04["张三"] = []string{"23", "18"}
	str04["李四"] = []int{2, 3}
	fmt.Println(str04)

}
