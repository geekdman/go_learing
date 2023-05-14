package main

import (
	"fmt"
	"os"
)

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
		var choice int
		fmt.Println("请输入您的选择：")
		fmt.Scan(&choice)

		switch choice {
		// 			   1. 查看客户
		case 1:
			fmt.Println("查看客户")
			fmt.Println("customers:::", customers)
			for _, customerMap := range customers {
				fmt.Printf("姓名：%v,年龄：%v,身高：%v\n", customerMap["name"], customerMap["age"], customerMap["height"])
			}

		//2. 添加客户
		case 2:
			fmt.Println("添加客户")
			var (
				name   string
				age    int
				height string
			)
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
			customers = append(customers, customer)

		//  3. 删除客户
		case 3:
			fmt.Println("删除客户")

		//  4. 修改客户
		case 4:
			fmt.Println("修改客户")
		case 5:
			fmt.Println("退出程序")
			os.Exit(200)
		default:
			fmt.Println("非法输入!")
		}

	}

}
