package main

import (
	"fmt"
	"os"
)

// 增加用户
func AddCoustomer(customers *[]map[string]interface{}, customer *map[string]interface{}) {
	*customers = append(*customers, *customer)
}

// 删除用户
func DeleteCoustomer(customers *[]map[string]interface{}, name string) {
	for i, customerMap := range *customers {
		if customerMap["name"] == name {
			fmt.Println()
			*customers = append((*customers)[0:i], (*customers)[i+1:]...)
		} else {
			fmt.Println("该用户不存在，请先添加用户")
		}
	}
}

func QueryCoustomer(customers *[]map[string]interface{}) {
	fmt.Println("查看客户")
	fmt.Println("customers:::", customers)
	for _, customerMap := range *customers {
		fmt.Printf("姓名：%v,年龄：%v,身高：%v\n", customerMap["name"], customerMap["age"], customerMap["height"])
	}
}
func UpdateCoustomer(customers *[]map[string]interface{}, customer *map[string]interface{}) {
	for i, customMap := range *customers {
		if customMap["name"] == (*customer)["name"] {
			(*customers)[i]["age"] = (*customer)["age"]
			(*customers)[i]["height"] = (*customer)["height"]
		} else {
			fmt.Println("该用户不存在，请先添加用户")
		}
	}
}

func main() {

	//var s = [][]string{{"yuan", "22", "180cm"}, {"yuan", "25", "180cm"}, {"yuan", "22", "180cm"}}
	//fmt.Println(s[1][1])

	// 使用切片，来实现数据的有序，map 来管理数据
	var customers []map[string]interface{}

	for {
		fmt.Println(`
			   1. 查看客户
			   2. 添加客户
			   3. 删除客户
			   4. 修改客户
			   5. 退出程序
		`)

		var (
			name   string
			age    int
			height string
			choice int
		)
		fmt.Println("请输入您的选择：")
		fmt.Scan(&choice)

		switch choice {
		// 			   1. 查看客户
		case 1:
			QueryCoustomer(&customers)
		//2. 添加客户
		case 2:
			fmt.Println("添加客户")
			customer := make(map[string]interface{})

			fmt.Println("请输入添加客户名字")
			fmt.Scan(&name)

			fmt.Println("请输入添加客户年龄")
			fmt.Scan(&age)

			fmt.Println("请输入添加客户身高")
			fmt.Scan(&height)
			// 创建一个新的客户map对象
			customer = map[string]interface{}{"name": name, "age": age, "height": height}
			// 将创建的map对象追加到customers
			AddCoustomer(&customers, &customer)

		//  3. 删除客户
		case 3:
			fmt.Println("请输入要删除客户的姓名")
			fmt.Scan(&name)
			DeleteCoustomer(&customers, name)

		//  4. 修改客户
		case 4:
			fmt.Println("请输入修改客户名字")
			fmt.Scan(&name)

			fmt.Println("请输入修改客户年龄")
			fmt.Scan(&age)

			fmt.Println("请输入修改客户身高")
			fmt.Scan(&height)
			customer := map[string]interface{}{"name": name, "age": age, "height": height}
			UpdateCoustomer(&customers, &customer)
		case 5:
			fmt.Println("退出程序")
			os.Exit(200)
		default:
			fmt.Println("非法输入!")
		}

	}
}
