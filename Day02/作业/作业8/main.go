package main

import (
	"fmt"
)

/*
基于切片实现客户关系管理系统(客户增删改查)。客户信息包括姓 名、年龄、职业。
数据结构采用
var customers = [][]string{[]string{"yuan","23","CEO"}, []string{"alex","43","网管"}}
*/

//func addPerson(customers [][]string, info string) {
//	var person []string
//	info = strings.Trim(info, " ")
//	person = strings.Split(info, " ")
//	customers = append(customers, person)
//}
//func changeinfo(customers [][]string, info string) {
//	var p []string
//	fmt.Scan(&info)
//	info = strings.Trim(info, " ")
//	p = strings.Split(info, " ")
//	for _, v := range customers {
//		if v[0] == p[0] {
//			v[1] = p[1]
//			v[2] = p[2]
//		} else {
//			fmt.Printf("这个%s 不存在，需要先添加该客户", p[0])
//		}
//	}
//}
//func deletePerson(customers [][]string, info string) {
//
//}
//func queryinfo(customers [][]string, info string) {
//	var name string
//	fmt.Scan(&info)
//	name = strings.Trim(info, " ")
//	for _, v := range customers {
//		if v[0] == name {
//			fmt.Printf("姓名: %s;年龄: %s;职位: %s", v[0], v[1], v[2])
//		} else {
//			fmt.Printf("这个%s 不存在，需要先添加该客户", name)
//		}
//	}
//}

func main() {
	var (
		name string
		age  string
		job  string
	)
	var (
		customers [][]string
		choice    int
	)
	//customers = make([][]string, 1)
	for true {
		fmt.Print(`
客户关系管理系统：
	1、添加客户信息
	2、删除客户信息
	3、修改客户信息
	4、查询客户信息
	5、打印客户信息列表
请输入对应数字
	`)
		var p = make([]string, 3)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("请输入客户信息，如：张三 23 CEO")
			fmt.Scanln(&name, &age, &job)
			p[0] = name
			p[1] = age
			p[2] = job
			customers = append(customers, p)
		case 2: //删除
			fmt.Println("请输入要删除的客户姓名如：张三")
			fmt.Scan(&name)
			for i, v := range customers {
				if v[0] == name {
					customers = append(customers[0:i], customers[i+1:]...)
				}
			}
		case 3: //修改客户信息
			fmt.Println("请输入要修改的客户姓名如：张三")
			fmt.Scan(&name)
			var p1 []string
			for i, v := range customers {
				if v[0] == name {
					p1 = v
					fmt.Println(`
						请输入要修改的内容：
							1、修改客户姓名
							2、修改客户年龄
							3、修改客户职业
 					`)
					var t int
					fmt.Scan(&t)
					switch t {
					case 1:
						fmt.Println("请输入要修改客户的姓名")
						fmt.Scan(&name)
						p1[0] = name
					case 2:
						fmt.Println("请输入要修改客户的年龄")
						fmt.Scan(&age)
						p1[1] = age
					case 3:
						fmt.Println("请输入要修改客户的职业")
						fmt.Scan(&job)
						p1[2] = job
					}
				}
				if i == len(customers)-1 {
					fmt.Printf("这个%s 不存在，需要先添加该客户", name)
				}
			}

		case 4: //查
			fmt.Println("请输入要修改的客户姓名，如：张三 ")
			fmt.Scan(&name)
			for i, v := range customers {
				if v[0] == name {
					fmt.Printf("[%-2d] 姓名: %s;年龄: %s;职位: %s", i, v[0], v[1], v[2])
				}
				if i == len(v)-1 {
					fmt.Printf("这个人%s不存在", name)
				}
			}

		case 5:
			for _, v := range customers {
				fmt.Printf("姓名: %-10s; 年龄: %-20s; 职位: %-20s\n", v[0], v[1], v[2])
			}
		}
	}

}
