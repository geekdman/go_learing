package main

import (
	"fmt"
)

type  Person struct {
	// name , agem, isMarried,socres 是成员变量
	name string
	age int
	isMarried bool
	socres []string
}

func NewPerson(name string,age int,isMarried bool, socres []string) Person{
	return Person{name,age,isMarried,socres}
}

func (p Person) listen()  {
	fmt.Printf("%s听课",p.name)
}

func main()  {
	p1 := NewPerson("张三",18,true,[]string{"english"})
	fmt.Println(p1)
	p1.listen()
	p2 := NewPerson("lisi",18,true,[]string{"english"})
	fmt.Println(p2)
	p2.listen()
}
