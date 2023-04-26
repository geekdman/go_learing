package main

import "fmt"

/*
new() 与 make() 的区别
*/

func main() {
	var (
		p *int
		a []int
	)
	p=new(int)       //new用来为指针创建初始化内存空间，防止出现空指针异常
	fmt.Println(p,*p)
	a=make([]int,5)   //make 主要是为了slice等类型初始化内存空间，防止出现空指针异常
	fmt.Println(a)
}