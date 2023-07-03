package main

import (
	"fmt"
	"runtime"
)

type  Person struct {
	// name , agem, isMarried,socres 是成员变量
	name string
	age int
	isMarried bool
	socres []string
}

func NewPerson(name string,age int,isMarried bool, socres []string) *Person{
	return &Person{name,age,isMarried,socres}
}

func main()  {
	p1 := NewPerson("张三",18,true,["English","math"])
	fmt.Println(p1)

}
