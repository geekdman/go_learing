package main

import (
	"fmt"
	"os"
)

var (
	customers []map[string]interface{}
	name      string
	age       int
	height    string
	choice    int
	id        int
)

// 增加用户
func AddCoustomer() {
	customer := make(map[string]interface{})

	fmt.Println("请输入客户名字")
	fmt.Scan(&name)

	fmt.Println("请输入客户年龄")
	fmt.Scan(&age)

	fmt.Println("请输入客户身高")
	fmt.Scan(&height)
	// 创建一个新的客户map对象
	customer = map[string]interface{}{"name": name, "age": age, "height": height}
	customers = append(customers, customer)
}

// 删除用户
func DeleteCoustomer() {
	fmt.Print("请输入要删除客户编号：")
	fmt.Scan(&id)
	for i, customerMap := range customers {
		if customerMap["id"] == id {
			fmt.Println()
			customers = append(customers[0:i], customers[i+1:]...)
		} else {
			fmt.Println("该用户不存在，请重新输入")
		}
	}
}

// 查询客户
func QueryCoustomer() {
	fmt.Print("查看客户：")
	fmt.Println("customers:::", customers)
	for _, customerMap := range customers {
		fmt.Println("-------------------------------------------------")
		fmt.Printf("姓名：%-5v,年龄：%-5v,身高：%-5v\n", customerMap["name"], customerMap["age"], customerMap["height"])
		fmt.Println("-------------------------------------------------")
	}
}

// 更新客户
func UpdateCoustomer() {
	fmt.Println("请输入客户编号：")
	fmt.Scan(&id)
	fmt.Println("请输入修改客户名字")
	fmt.Scan(&name)

	fmt.Println("请输入修改客户年龄")
	fmt.Scan(&age)

	fmt.Println("请输入修改客户身高")
	fmt.Scan(&height)
	customer := map[string]interface{}{"name": name, "age": age, "height": height}
	for i, customMap := range customers {
		if customMap["name"] == customer["name"] {
			customers[i]["age"] = customer["age"]
			customers[i]["height"] = customer["height"]
		} else {
			fmt.Println("该用户不存在，请先添加用户")
		}
	}
}

func main() {

	//var s = [][]string{{"yuan", "22", "180cm"}, {"yuan", "25", "180cm"}, {"yuan", "22", "180cm"}}
	//fmt.Println(s[1][1])

	// 使用切片，来实现数据的有序，map 来管理数据

	for {
		fmt.Println(`
-------------------------
	   1. 查看客户
	   2. 添加客户
	   3. 删除客户
	   4. 修改客户
	   5. 退出程序
-------------------------
		`)

		fmt.Print("请输入您的选择：")
		fmt.Scan(&choice)

		switch choice {
		//  1. 查看客户
		case 1:
			QueryCoustomer()
		//  2. 添加客户
		case 2:
			AddCoustomer()
		//  3. 删除客户
		case 3:
			DeleteCoustomer()
		//  4. 修改客户
		case 4:
			UpdateCoustomer()
		case 5:
			fmt.Println("退出程序")
			os.Exit(200)
		default:
			fmt.Println("非法输入!")
		}

	}
}
