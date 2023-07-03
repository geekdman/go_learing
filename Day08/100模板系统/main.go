package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(context *gin.Context){
	var user = "Tony"
	context.HTML(http.StatusOK,"index.html",gin.H{
		"username": user,
		"state" : "在线",
		"booksSlice" : []string{},


	})
}

type Student struct {
	name string

}

func main() {

	r := gin.Default()

	//加载html 目录
	r.LoadHTMLGlob("temps/*")
	r.Static("/static","./static")

    r.GET("/index", index)

	r.Run("127.0.0.1:8080")

}
