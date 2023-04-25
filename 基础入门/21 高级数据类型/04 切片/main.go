package main

import "fmt"

func main() {
	// 切片的构建

	// 方式一： 对数组切片得到切片对象
	var arr = [5]int{11, 12, 13, 14, 15}
	s1 := arr[0:3]
	s2 := arr[1:4]
	fmt.Println(s1, s2)
	fmt.Printf("%T,%T,%T", arr, s1, s2) //[5]int (数组类型),[]int：切片类型,[]int： 切片类型

	arr[2] = 100

	//切片的存储原理
	//
	var a = [...]int{1, 2, 3, 4, 5, 6}
	a1 := a[0:3]
	a2 := a[0:5]
	a3 := a[1:5]
	a4 := a[1:]
	a5 := a[:]
	a6 := a3[1:2]
	fmt.Printf("a1的长度%d,容量%d\n", len(a1), cap(a1))
	fmt.Printf("a2的长度%d,容量%d\n", len(a2), cap(a2))
	fmt.Printf("a3的长度%d,容量%d\n", len(a3), cap(a3))
	fmt.Printf("a4的长度%d,容量%d\n", len(a4), cap(a4))
	fmt.Printf("a5的长度%d,容量%d\n", len(a5), cap(a5))
	fmt.Printf("a6的长度%d,容量%d\n", len(a6), cap(a6))

	// 方式二：直接定义切片
	// 其中切片可以从数组和切片中获取，也可以直接声明一个切片，每一种数据类型都可以拥有其切片类型，表示多个相同类型元素的连续集合。因此切片类型也可以被声明
	// 切片类型 声明格式如下：
	// var name []Type // []Type 是切片类型的标识
	var name []string
	fmt.Println(name)

	// 方式三：声明并且初始化
	var s = []int{1, 2, 3, 3, 4}
	fmt.Println(s)
	s3 := s[1:]
	s3[1] = 4
	fmt.Println(s3, "\n", s)

	// 练习：考察值拷贝和 slice 的结构和原理
	var arr1 = []int{1, 2, 3}
	arr2 := arr1
	arr1[0] = 100
	fmt.Println(arr2)
}
