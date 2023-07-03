package main

import (
	"github.com/gin-gonic/gin"
)

/*
  路由分组
*/

// 定义路由函数
func  GetBook(c  *gin.Context)  {
	c.JSON(200,gin.H{
		"data":"查看书籍",
	})
}
func  AddBook(c  *gin.Context)  {
	c.JSON(200,gin.H{
		"data":"添加书籍",
	})
}
func  EditBook(c  *gin.Context)  {
	c.JSON(200,gin.H{
		"data":"更新书籍信息",
	})
}
func  DeleteBook(c  *gin.Context)  {
	c.JSON(200,gin.H{
		"data":"删除书籍",
	})
}


// 主函数
func main() {

	r := gin.Default()

	// 书籍相关的路由
	// http://localhost:8080/books/
	bookRouter:= r.Group("books")
	{
		// http://localhost:8080/books/
		bookRouter.GET("/", GetBook)
		// http://localhost:8080/books/add
		bookRouter.POST("/add", AddBook)
		// http://localhost:8080/books/edit
		bookRouter.PUT("/edit", EditBook)
		// http://localhost:8080/books/delete
		bookRouter.DELETE("/delete", DeleteBook)
	}

	// 路径匹配失败,执行操作
	r.NoRoute(func(context *gin.Context) {
		context.JSON(200,gin.H{
			"code": 1001,
			"msg": "路由匹配失败",
		})
	})


	// 启动服务
	r.Run("127.0.0.1:8000")
}