package main

import (
	"fmt"
)

/*
1、位置参数
   func add(x int, y int) {}
   func add(x, y int) {}
2、不定长参数
   func add(nums ...int) {}
    备注：系统会把num 定义为一个切片
*/

// 位置参数
// func add(x int, y int) {} // 下面是简写参数
func add(x, y int) {
	fmt.Println(x + y)
}

// 不定长参数
func add2(nums ...int) {
	ret := 0
	for _, num := range nums {
		ret += num
	}
	fmt.Println(ret)
}

// 案例：
// 	 位置参数+ 不定长参数
//  当位置参数和不定长参数混合使用的时候，位置参数要放到不定长参数之前。

func cal(oper string, nums ...int) {
	var ret int
	switch oper {
	case "+":
		ret = 0
		for _, num := range nums {
			ret += num
		}
	case "*":
		ret = 1
		for _, num := range nums {
			ret *= num
		}
	}
	fmt.Println(ret)
}

func main() {
	add(10, 20)
	add2(5, 10, 20, 30)
	cal("*", 1, 2, 3, 4)
	cal("+", 1, 2, 3, 4)
}
