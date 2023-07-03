package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
1、获取基本请求数据
2、获取路径参数
    url 路径：/user/:id
    常用函数：
	func(c*Context)Param(keystring)string
3、获取查询参数
    url 路径：/path?id=1234&name=Manu&value=111

    常用函数：
	func(c*Context)Query(keystring)string
	func(c*Context)DefaultQuery(key,defaultValuestring)string
	func(c*Context)GetQuery(keystring)(string,bool)

5、获取PostForm数据
	客户端请求类型为form-data 、x-www-form-urlencoded
	func(c*Context)PostForm(keystring)string
	func(c*Context)DefaultPostForm(key,defaultValuestring)string
	func(c*Context)GetPostForm(keystring)(string,bool)
4、ShouldBind函数

*/

//1、获取基本请求数据
func getRequest(context *gin.Context)  {

	fmt.Println(context.Request)
	fmt.Println(context.Request.Method)
	fmt.Println(context.Request.URL)
	fmt.Println(context.Request.RemoteAddr)
	fmt.Println(context.Request.Header)
	fmt.Println(context.Request.Header["User-Agent"])   // 获取User-Agent
	fmt.Println(context.GetHeader("User-Agent"))   //获取User-Agent
}

type Book struct{
	Name string `json:"name" form:"name"`
	Title string  `json:"title" form:"title"`
}

func main() {
	r :=gin.Default()

	//r.GET("/book/:username/id/:id", func(context *gin.Context) {
	//	//username := context.Param("username")
	//	//id := context.Param("id")
	//	context.JSON(200,gin.H{
	//		"username": context.Param("username"),
	//		"id": context.Param("id"),
	//	})
	//
	//})

	r.GET("/test", func(context *gin.Context) {
		//context.DefaultQuery("page","1")
		context.JSON(200,gin.H{
			"page": context.DefaultQuery("page","1"),
		})
	})
	r.POST("/book", func(context *gin.Context) {
		book := Book{}
		if context.ShouldBind(&book) != nil {
			fmt.Println("解析失败")

		}
		fmt.Println("book:::",book)
		context.JSON(200,gin.H{
			"book": book,
		})
	})
    // 启动服务
	r.Run()
}
