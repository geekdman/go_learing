package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type BaseModel struct {
	ID         int       `gorm:"primaryKey"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
}
// 教师表
type Teacher struct {
	BaseModel
	Tno    int
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
}
//教室
type Class struct {
	BaseModel
	Num     int
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	TeacherID int
	Teacher   Teacher
}
type Course struct {
	BaseModel
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	Credit int
	Period int

	// 多对一
	TeacherID int
	Teacher   Teacher
}
type  Student struct {
	BaseModel
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	Sno int
	Pwd string `gorm:"type:varchar(100);not null"`
    Gender byte   `gorm:"default:1"`
    Birth  *time.Time
    Remark string `gorm:"type:varchar(255);"`

    // 多对一
    ClassID int
    Class   Class `gorm:"foreignKey:ClassID"`
    // 多对多
    Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

func addRecord(db *gorm.DB)  {
	// 创建老师
	t1 := Teacher{BaseModel: BaseModel{ID: 1}, Name:"Tony",Tno: 1001, Pwd: "123"}
	db.Create(&t1)
	fmt.Println(t1.ID)
	t2 := Teacher{BaseModel: BaseModel{ID: 2}, Name:"Tom",Tno: 1002, Pwd: "234"}
	db.Create(&t2)
	t3 := Teacher{BaseModel: BaseModel{ID: 3},Name:"Rain",Tno: 1002, Pwd: "345"}
	db.Create(&t3)


	// 创建班级
	c1:=Class{Name:"软件1班",Num: 78,TeacherID: 1}
	c2:=Class{Name:"软件2班",Num: 88,TeacherID: 2}
	c3:=Class{Name:"软件3班",Num: 90,TeacherID: 3}
	c4:=Class{Name:"软件4班",Num: 32,TeacherID: 1}
	classes  := []Class{c1,c2,c3,c4}
	db.Create(&classes)

	// 创建课程
	course01 := Course{Name: "计算机网络", Credit: 3, Period: 16, TeacherID: 1}
	course02 := Course{Name: "数据结构", Credit: 2, Period: 24, TeacherID: 1}
	course03 := Course{Name: "数据库", Credit: 2, Period: 16, TeacherID: 2}
	course04 := Course{Name: "数字电路", Credit: 3, Period: 12, TeacherID: 2}
	course05 := Course{Name: "模拟电路", Credit: 1, Period: 8, TeacherID: 2}
	//批量创建

	courses := []Course{course01,course02,course03,course04,course05}
	db.Create(&courses)

}
func selectRecord(db *gorm.DB) {
	// (1) 查询全部记录
	/*var teachers []Teacher
	db.Find(&teachers)
	fmt.Println("teachers:", teachers)

	for i, v := range teachers {
		fmt.Println(i, v.Name, v.Tno)
	}*/

	// (2) 查询单条记录
	/*course := Course{}
	db.Take(&course)
	db.First(&course)
	db.Last(&course)
	db.Where("credit < ?", 3).Order("credit").Take(&course)
	fmt.Println("course", course)*/

	// (3) Where语句

	/*
		// 字符串Where语句
		var courses []Course
		db.Where("credit > ?", 1).Find(&courses)
		db.Where("credit between ? and ?", 1, 3).Find(&courses)
		var total int64
		db.Model(&Course{}).Where("credit between ? and ?", 1, 2).Count(&total)
		fmt.Println("total:", total)*/

	/*var courses []Course
	db.Where("credit >? and period<?", 1, 20).Find(&courses)
	fmt.Println("courses:", courses)
	// 结构体Where语句
	db.Where(Course{Credit: 0, Period: 12}).Find(&courses)
	fmt.Println("courses:", courses)
	// map的Where语句
	db.Where(map[string]interface{}{"Credit": 0, "Period": 12}).Find(&courses)
	fmt.Println("courses:", courses)*/
}
func updateRecord()  {
	
}
func deleteRecord()  {
	
}

func main() {

	dsn := "root:password@tcp(10.0.11.7:3306)/slectclass?charset=utf8mb4&parseTime=True&loc=Local"


	//
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel: logger.Info, // Log level
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 日志配置
	})
	if err != nil {
		//panic("failed to connect database")
		log.Fatalf("xxxxxxxxxxxx")
	}
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Student{})
	//addRecord(db)
	//selectRecord(db)
}