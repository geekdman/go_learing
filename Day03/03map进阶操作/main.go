package main

import "fmt"

/*
map 进阶操作

	一、	切片嵌套map
	二、 map 嵌套map
	三、 map 嵌套切片
*/
func main() {
	//一、切片嵌套map
	/*
		var stu01 = map[string]string{"name": "user01", "age": "32"}
		var stu02 = map[string]string{"name": "user02", "age": "18"}
		var stu03 = map[string]string{"name": "user03", "age": "20"}

		// 切片管理
		//var stus = []map[string]string{{"name": "tom", "age": "32"}, {"name": "user02", "age": "32"}, {"name": "user03", "age": "32"}}
		var stus = []map[string]string{stu01, stu02, stu03}

		//fmt.Println(stus)
		// 1、查询第二个学生的年龄
		fmt.Println(stus)
		fmt.Println(stus[2])
		fmt.Println(stus[2]["age"])

		// 2、查询每个学生的信息
		for i, stuMap := range stus {
			fmt.Printf("%d----姓名：%s 年龄：%s\n", i, stuMap["name"], stuMap["age"])
		}

		// 3、添加学生信息
		var (
			name string
			age  string
		)
		fmt.Println("请输入姓名")
		fmt.Scan(&name)

		fmt.Println("请输入年龄")
		fmt.Scan(&age)

		var stu = map[string]string{"name": name, "age": age}
		stus = append(stus, stu)
		fmt.Println(stus)

		var a1 = map[string][]string{"user01": {"tom", "32", "18cm"}}
		fmt.Println(reflect.TypeOf(a1))
		fmt.Println(a1)*/

	//	二、map嵌套map
	/*	var stu01 = map[string]string{"name": "user01", "age": "32"}
		var stu02 = map[string]string{"name": "user02", "age": "18"}
		var stu03 = map[string]string{"name": "user03", "age": "20"}
		var stus = map[int]map[string]string{1001: stu01, 1002: stu02, 1003: stu03}

		// 1、查询1002 学生的姓名
		fmt.Printf("(1) 1002 学号的学生的姓名为 %s", stus[1002]["name"])
	*/
	// 三、map嵌套切片
	var stu01 = map[string]interface{}{"name": "user01", "age": 32, "hobby": []string{"抽烟", "喝酒", "烫头"}}

	var stus = []map[string]interface{}{stu01}
	fmt.Println(stus[0]["hobby"].([]string)[2]) //([]string)[2], 断言，把类型给系统
	fmt.Println(stus[0]["hobby"].([]string)[2]) //([]string)[2], 断言，把类型给系统

}
