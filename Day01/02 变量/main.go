package main

import "fmt"

/*
作者： dman
标题： 变量
内容：
	1、 变量声明，格式如下：
		标准定义
			var name type = expression
		其中 type 和 expression 可以省略其中一个，但是不能同时省略
			var name type            常用于变量的批量定义,或者结构体的定义
			var name = expression    常用于常量定义
        短变量声明
			name := expression       常用于接受函数返回值,或者局部变量
    2、变量赋值
       变量名 = 值
       变量名 = 值+ 值
       变量名 = 变量名
	3、匿名变量
		匿名变量 就是没有命名的变量,在使用多重赋值的时候,如果想要忽略某个值, 可以使用匿名变量, 匿名变量使用一个下划线表示。
		用途： 常用在接受函数返回值的场景
	4、变量命名规则

变量注意事项:

	1、Go语言中定义的变量，必须被调用，不然编译的时候会报错
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
	//1、 变量声明
	var v1 int = 100 // 标准声明，var name type = expression
	fmt.Println(v1)

	var v2 = 100 // 省略 type 声明,
	fmt.Println(v2)

	var v3 int // 省略expression 声明,
	fmt.Println(v3)

	v4 := 100
	fmt.Println(v4) // 短变量声明
	//2、匿名变量

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
