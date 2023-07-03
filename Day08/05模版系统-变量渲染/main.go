package main

import "github.com/gin-gonic/gin"

type Student struct {
	Name string
	Age  int
}

func main() {
	// 基于获取引擎对象，可以理解为路由对象
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(context *gin.Context) {

		context.HTML(200, "index.html", gin.H{
			"user":       "yuan",
			"state":      "在线",
			"booksSlice": []string{"金瓶梅", "聊斋", "剪灯新话", "国色天香"},
			"stuMap": map[string]interface{}{
				"name":  "rain",
				"age":   22,
				"hobby": []string{"足球", "篮球", "双色球"},
			},
			"stuStruct": Student{Name: "yuan", Age: 22},
		})
	})

	// 启动:默认本机的8080端口
	r.Run()
}