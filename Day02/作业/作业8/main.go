package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync/atomic"
)

/*
	基于切片实现客户关系管理系统(客户增删改查)。客户信息包括姓 名、年龄、职业。
	数据结构采用
	var customers = [][]string{[]string{"yuan","23","CEO"}, []string{"alex","43","网管"}}
*/
func main()  {
	var (
		customers [][]string
		customer []string
		choice int
	)
	fmt.Print(`
	客户关系管理系统：
		1、添加客户信息
    	2、删除客户信息
   		3、修改客户信息
		4、查询客户信息
`)
	fmt.Println(customers)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println()
	case 2:

	case 3:
	case 4:

	}
	s:=[]int{1,2,3,4,5}
	sort.Reverse(sort.IntSlice{s})
	reflect.Swapper()
}