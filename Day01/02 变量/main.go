package main

import "fmt"

/*
	1、go定义的变量必须被调用，不然会报错
		报错如下：
		// .\main.go:4:6: x declared and not used

	2、变量的语法：
		先声明，再赋值
		声明并赋值
		简写形式
		多变量赋值
	3、拓展
		变量不能重复声明
		变量赋值
		值拷贝，当一个变量赋值给另外一个变量的时候，会发生值拷贝。

	4、匿名变量 (_)
		var _,b = 10,20
		用途，取函数返回值的时候使用，只获取我们想要的变量，节省内存 var _,X= foo()
	5、变量的命名
		见名知意
        数字不能开头，
		下划线和横杠可以作为变量名，可以下划线开头
        @！# 等特殊不能作为变量名

*/

func main() {

	var x = 100
	fmt.Println(x)
	var y = 200

	// 2、变量语法
	fmt.Println(x + y)
	// 先声明，再赋值
	var a int8
	a = 100
	fmt.Println(a)
	// 声明并赋值, 编译器会根据后面的值去判断变量类型。
	var b = 100
	fmt.Println(b)
	// 简写形式
	c := 300 // 等同于 var c = 300
	fmt.Println(c)
	// 多变量赋值
	var d, e, f = 10, 20, 30
	fmt.Println(d, e, f)

	//3、拓展
	// 变量不能重复声明
	var x = 100
	var x = 200
	//案例2
	var r, s = 10, 20
	var z = r + s // 必须使用var ,或者 z := r + s
	fmt.Println(z)
	//案例3 ，值拷贝
	var o = 100
	var i = o //发生值拷贝,
	fmt.Println(i)

	//	4、匿名变量
	var _, a1 = 10, 20
	fmt.Println(a1)
}
