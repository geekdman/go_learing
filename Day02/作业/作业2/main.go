package main

import (
	"fmt"
)

/*
解析执行过程(描述过程)
	p1 := 1
	p2 := &p1
	*p2++
	fmt.Println(p1)
	fmt.Println(*p2)
*/


func main() {
	p1 := 1            // 声明并赋值 p1 变量为int 类型，初始化值为1
	p2 := &p1  			// 声明p2（为指针类型）,并将p1 的地址赋给p2。将p2 指向p1
	//fmt.Println(reflect.TypeOf(p2))
	*p2++               //将p2 所指向的内存空间存放的int类型的 1 进行+1 操作，得到2。此时，由于 p2 是指向p1 的内存地址，所以，此时p1 = 2
	fmt.Println(p1)    // 2
	fmt.Println(*p2)   //2
}
