package main

import (
	"encoding/json"
)

/*
	序列化
序列化： 通过某种方式把数据结构或对象写入到磁盘文件中或通过网络传到其他节点的过程

反序列化：把磁盘中对象或者把网络节点中传输的数据恢复为数据对象的过程。
*/

func main() {
	//var stu01 = map[string]string{"name": "user01", "age": "32"}
	//var stu02 = map[string]string{"name": "user02", "age": "18"}
	//var stu03 = map[string]string{"name": "user03", "age": "20"}

	/*	//切片管理
		var stus = []map[string]string{stu01, stu02, stu03}

		// json序列化
		ret, _ := json.Marshal(stus)
		fmt.Println(string(ret))
	*/
	//  json 反序列化
	s := `{"name":"tom","age":"18"}`
	var data []map[string]string
	json.Unmarshal([]byte(s), &data)

}
