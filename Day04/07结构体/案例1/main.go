package main

type Student struct {
	sid int64
	name string
	age int
	Course []string
}

func CourseInt(stu Student)  {
	stu.Course = []string{"chinese","math"}
}

func CourseInt2(stu *Student){
	(*stu).Course = []string{"chinese","math"}
}

func main()  {
	//案例1
	s1:= Student{sid:""}

}