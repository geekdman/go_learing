package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
1、s:= "hello golang",将字符串s中golang 字符串切片出来
2、var x = 10; var y= 5 , 将y 的值 通过三种方式加到x中？
3、程序执行结果
	var a int8
*/

func main() {
	/*
		1、s:= "hello golang",将字符串s中golang 字符串切片出来
	*/
	s := "hello golang"
	s_index := strings.Index(s, "golang")
	fmt.Println(string(s[s_index:]))

	/*
		2、var x = 10; var y= 5 , 将y 的值 通过三种方式加到x中？
	*/
	var x = 10
	var y = 5
	//x += y
	x = x + y

	/*
		3、程序执行结果
		var a int8 = 100
		b := a
		a++
		fmt.Println("a:", a)
		fmt.Println("b:", b)
	*/
	var a int8 = 100
	b := a
	a++
	fmt.Println("a:", a) // a = 101
	fmt.Println("b:", b) // b = 100

	/*
		4、引导用户输入生日字符串，格式为`“年-月-日”，
		比如`“1990-3-16”`，然后按`“您的生日是1990年-3月-16日”`的格式化字符串输出到终端控制台（**暂放**）。
	*/
	var birth string
	//var s_birth []string
	fmt.Println("请输入生日字符串，格式为'年-月-日',eg: 1994-01-05: ")
	fmt.Scan(&birth)
	birth = strings.Trim(birth, " ")
	var s_birth []string = strings.Split(birth, "-")
	fmt.Printf("您的生日是 %s年-%s月-%s日\n", s_birth[0], s_birth[1], s_birth[2])

	/*
		5、引导用户输入计算表达式，构建一个双位计算器功能，比如用户输入 1+2 或 3*4 或 4/2 或 5-8等关于加减乘除运算，打印最终的计算结果。`（基于switch语句）
	*/
	var (
		exp string
		ret []string
	)
	fmt.Println()
	fmt.Scan(&exp)

	if strings.Contains(exp, "+") {
		ret, _ = strings.Split(exp, "+")
	} else if strings.Contains(exp, "-") {
		ret, _ = strings.Split(exp, "-")
	} else if strings.Split(exp, "*") {
		ret, _ = strings.Split(exp, "*")
	} else if strings.Contains(exp, "/") {
		ret, _ = strings.Split(exp, "/")
	}
	exp1, _ := strconv.Atoi(ret[0])
	exp2, _ := strconv.Atoi(ret[0])
	fmt.Printf(exp1)
	/*
		6、程序随机内置一个位于一定范围内的数字作为猜测的结果，由用户猜测此数字。用户每猜测一次，由系统提示猜测结果：太大了、太小了或者猜对了，直到用户猜对结果或者猜测次数用完导致失败。

		1. 比如数字是1-100之间，设定一个数字比如：66。
		2. 让用户三次机会猜数字，如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了;只有等于66，显示猜测结果正确，退出循环。
		3. 最多三次都没有猜测正确，退出循环，并显示‘都没猜对,继续努力’。
	*/
	var (
		input1 int
	)
	isPass := false
	for i := 0; i < 3; i++ {
		rand_num := rand.Intn(99) + 1
		fmt.Scan(&input1)
		if input1 > rand_num {
			fmt.Println("猜测的结果大了")
		} else if input1 == rand_num {
			fmt.Printf("猜测结果正确,正确的答案是 %d", rand_num)
			isPass = true
			break
		} else if input1 < rand_num {
			fmt.Println("猜测的结果小了")
		}
		if isPass {
			fmt.Println("恭喜")
		} else {
			fmt.Println("三次机会使用完，都没猜对,继续努力")
		}
	}

}
