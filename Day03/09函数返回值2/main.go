package main

/*
函数返回值
格式：
1、一个返回值
func 函数名(参数名 参数类型) 返回值类型 {
	函数体
	retrun value
}
2、函数返回多个值
func 函数名(参数名 参数类型) (返回值1类型,返回值2类型) {
	函数体
	retrun value1,value2
}

3、函数返回值命名，可选

func 函数名(参数名 参数类型) (返回值变量1 返回值1类型) {
	函数体
	retrun
}
func 函数名(参数名 参数类型) (返回值变量1 返回值1类型, 返回值变量 返回值2类型) {
	函数体
	retrun
}
*/

func add(x, y int) int {
	return x + y
}

func add2(x, y int) (int, int) {
	ret1 := x + y
	return ret1, 200
}

func bar() int {
	ret := 100
	return ret
}

// 函数返回值的几种用法
func bar2() (ret int) {
	ret = 100
	return ret
}

func bar3() (ret int) {
	ret = 100
	return
}

func bar4() (ret, sum int) {
	ret = 100
	sum = ret + 100
	return
}

func main() {
	/*	// 案例1：
		var ret1 = add(3, 4)
		var ret2 = add(4, 5)
		fmt.Println(ret1)
		fmt.Println(ret2)*/

	/*	// 案例2
		ret3, ret4 := add2(5, 10)
		fmt.Println(ret3, ret4)
	*/

}
