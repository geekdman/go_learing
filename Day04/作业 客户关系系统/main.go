package main

import (
"fmt"
"os"
)

type Customer struct {
	Name      string 	`json:"name"`
	Age       int
	Height    string
	//Courses   []string
}

type CustomerManage struct {
	customers []Customer
}

// 增加用户
func (cm *CustomerManage)AddCoustomer() {
	var (
		Name      string
		Age       int
		Height    string
	)

	fmt.Println("请输入客户名字")
	fmt.Scan(Name)

	fmt.Println("请输入客户年龄")
	fmt.Scan(Age)

	fmt.Println("请输入客户身高")
	fmt.Scan(Height)
	//fmt.Println("请输入客户报名课程")
	//fmt.Scan(c.Courses)
	// 创建一个新的客户map对象
	c := Customer{Name,Age,Height}
	cm = append((*)cm, c)
}

// 删除用户
func (cm *CustomerManage) DeleteCoustomer() {
	if len(cm.customers) == 0 {
		fmt.Println("请先添加拥护")
		return
	}
	var id int
	fmt.Print("请输入要删除客户编号：")
	fmt.Scan(&id)
	for i, _:= range cm.customers {
		if i == id {
			fmt.Println()
			cm.customers = append(cm.customers[0:i], cm.customers[i+1:]...)
		} else {
			fmt.Println("该用户不存在，请重新输入")
		}
	}
}

// 查询客户
func (cm *CustomerManage) QueryCoustomer() {
	var id int
	fmt.Print("请输入要删除客户编号：")
	fmt.Scan(&id)
	for i, v := range cm.customers {
		if i == id {
			fmt.Println("-------------------------------------------------")
			fmt.Printf("姓名：%-5v,年龄：%-5v,身高：%-5v\n",v.Name, v.Age, v.Height)
			fmt.Println("-------------------------------------------------")
		}
	}
}

// 更新客户
func (cm *CustomerManage) UpdateCoustomer() {
	var id int
	fmt.Println("请输入客户编号：")
	fmt.Scan(&id)

	for i, v := range cm.customers {
		if id == i {
			fmt.Println("请输入修改客户名字")
			fmt.Scan(v.Name)
			fmt.Println("请输入修改客户年龄")
			fmt.Scan(v.Age)
			fmt.Println("请输入修改客户身高")
			fmt.Scan(v.Height)
		}

	}
}

func main() {

	//var s = [][]string{{"yuan", "22", "180cm"}, {"yuan", "25", "180cm"}, {"yuan", "22", "180cm"}}
	//fmt.Println(s[1][1])

	// 使用切片，来实现数据的有序，map 来管理数据

	cm
	var choice int
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
			cm.QueryCoustomer()
		//  2. 添加客户
		case 2:
			cm.AddCoustomer()
		//  3. 删除客户
		case 3:
			cm.DeleteCoustomer()
		//  4. 修改客户
		case 4:
			cm.UpdateCoustomer()
		case 5:
			fmt.Println("退出程序")
			os.Exit(200)
		default:
			fmt.Println("非法输入!")
		}

	}
}


