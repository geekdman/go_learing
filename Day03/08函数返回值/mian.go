package main

import (
	"fmt"
)

/*
函数返回值
格式：
func 函数名(参数名 参数类型) 返回值类型 {
	函数体
	retrun
}

func add(x, y int) int {}
*/

func add(x, y int) int {
	return x + y
}

// 不定长参数
func add2(nums ...int) {
	ret := 0
	for _, num := range nums {
		ret += num
	}
	fmt.Println(ret)
}

// 位置参数和不定长参数混合使用
func cal(oper string, nums ...int) int {
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
	return ret
}

func main() {
	// 案例1：
	var ret1 = add(3, 4)
	var ret2 = add(4, 5)
	fmt.Println(ret1)
	fmt.Println(ret2)

	var ret3 = cal("*", ret1, ret2)
	fmt.Println(ret3)

}
