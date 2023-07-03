package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)
/*

1、响应简单字符串
2、响应html页面
3、响应json
4、响应xml数据
5、响应protobuf数据

*/

func main() {
	r :=gin.Default()

	//1、响应简单字符串
	r.GET("/str", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	//2、响应html页面
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html",nil)
	})

	//3、响应json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "json数据",
		})
	})
	//4、响应xml数据
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "xml消息"})
	})

	//5、响应protobuf数据

	r.GET("/protobu", func(c *gin.Context) {
		reps := []int64{int64(9),int64(55)}

		// 自定义数据
		label := "你好啊,今天天气挺好的"

		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		fmt.Println(data)
		c.ProtoBuf(200,data)
	})

	// 重定向
	r.NoRoute(func(context *gin.Context) {
		//context.Redirect(http.StatusMovedPermanently,"/xml")
		context.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")

	})

	// 异步响应
	r.GET("/async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()

		c.JSON(200, gin.H{"msg": "异步执行中"})

	})
	// 2.同步
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
		c.JSON(200, gin.H{"msg": "同步执行结束"})
	})

	// 启动服务
	r.Run()
}
