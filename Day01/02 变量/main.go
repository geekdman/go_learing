package main

import "fmt"

/*
作者： dman
标题： 变量
内容：
	一、 变量声明，格式如下：
		标准定义
			var name type = expression
		其中 type 和 expression 可以省略其中一个，但是不能同时省略
			var name type            常用于变量的批量定义,或者结构体的定义
			var name = expression    常用于常量定义
        短变量声明
			name := expression       常用于接受函数返回值,或者局部变量
    二、变量赋值
       变量名 = 值
       变量名 = 值+ 值
       变量名 = 变量名

	三、匿名变量
		匿名变量 就是没有命名的变量,在使用多重赋值的时候,如果想要忽略某个值, 可以使用匿名变量, 匿名变量使用一个下划线表示。
		用途： 常用在接受函数返回值的场景
	四、变量命名规则
		1、变量名称必须由数字`、`字母、下划线组成。
		2、标识符开头不能是数字。
		3、标识符不能是保留字和关键字。
		4、建议使用驼峰式命名，当名字有几个单词组成的时优先使用大小写分隔
		5、变量名尽量做到见名知意。
		6、变量命名区分大小写

	五、多重变量赋值
       var a,b = 3 ,4

注意事项:

	1、Go语言中定义的变量，必须被调用，不然编译的时候会报错
		报错如下：
		// .\main.go:4:6: x declared and not used

	2、拓展
		变量不能重复声明
		变量赋值
		值拷贝，当一个变量赋值给另外一个变量的时候，会发生值拷贝。

	3、匿名变量 (_)
		var _,b = 10,20
		用途，取函数返回值的时候使用，只获取我们想要的变量，节省内存 var _,X= foo() ,可以节省内存空间。
	5、变量的命名
		见名知意
        数字不能开头，
		下划线和横杠可以作为变量名，可以下划线开头
        @！# 等特殊不能作为变量名
	6、通过变量名赋值给另外的变量，之间进行的是值拷贝，如下：
		var a =10 ; b := a
       其中b 和a 在内存中的地址完全不同，当执行b := a 的时候，系统会将a 的数据在内存中拷贝一份，b指代的就是副本的内容。这样修改b 并不会影响a。这样也就出现了一个Go语言的基本概念。所有拷贝都是值拷贝，没有所谓的引用拷贝

*/

func main() {
	// 一、 变量声明
	var v1 int = 100 // 标准声明，var name type = expression
	fmt.Println(v1)

	var v2 = 100 // 省略 type 声明,
	fmt.Println(v2)

	var v3 int // 省略expression 声明,
	fmt.Println(v3)

	v4 := 100
	fmt.Println(v4) // 短变量声明
	// 二、变量赋值
	var v5 int
	v5 = 10 //变量名=值
	fmt.Println(v5)

	var s1 string = "hello world" //变量名=值
	fmt.Println(s1)

	s2 := "hello nihao"
	fmt.Println(s2)

	var name, age = "tony", 18 //变量名=值
	fmt.Println(name, age)

	var n1, n2 = 10, 20
	var n3 = n1 + n2 //变量名= 值 +值
	fmt.Println(n3)

	var n4 = 100
	var n5 = n4
	fmt.Println(n5) //变量名 = 变量名（值拷贝）

	// 三、匿名变量
	var n6, n7, n8 = 4, 5, 6
	fmt.Println(n6, n7, n8)
	_, _, n10 := n6, n7, n8
	fmt.Println(n10)

	// 四、变量命名规则
	var arr, _arr, arr4 string
	fmt.Println(arr, _arr, arr4)

	var execCmd string
	fmt.Println(execCmd)

	//五、拓展
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
	
}
