package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
type User struct {
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Email string `json:"email" form:"email"`
}

func main() {
	r :=gin.Default()

/*	// 请求基本数据
	// http://127.0.0.1:8080/test
	r.GET("/test", func(context *gin.Context) {
		fmt.Println(context.Request)
		fmt.Println(context.Request.Method)
		fmt.Println(context.Request.URL)
		fmt.Println(context.Request.RemoteAddr)
		fmt.Println(context.Request.Header)
		fmt.Println(context.Request.Header["User-Agent"])   // 获取User-Agent
		fmt.Println(context.GetHeader("User-Agent"))   //获取User-Agent
	})*/

/*    // 获取路径参数
    // http://127.0.0.1:8080/test/user01/id/1001
	r.GET("/test/:username/id/:id", func(context *gin.Context) {
		//username := context.Param("username")
		//id := context.Param("id")
		context.JSON(200,gin.H{
			"username": context.Param("username"),
			"id": context.Param("id"),
		})

	})*/

/*	// 获取查询参数
	// http://127.0.0.1:8080/test?page=3
	r.GET("/test", func(context *gin.Context) {
		//context.DefaultQuery("page","1")
		context.JSON(200,gin.H{
			"page": context.DefaultQuery("page","1"),
		})
	})
	*/

/*    //获取表单数据，类型为form-data，或者x-www-form-urlencoded
	r.POST("test", func(c *gin.Context) {
		//获取name参数,通过PostForm获取的参数值是String类型。
		name:=c.PostForm("name")
		//跟PostForm的区别是可以通过第⼆个参数设置参数默认值

		//name:=c.DefaultPostForm("name","root")
		//获取id参数,通过GetPostForm获取的参数值也是String类型,
		//获取id参数,通过GetPostForm获取的参数值也是String类型,
		//区别是GetPostForm返回两个参数，第⼀个是参数值，第⼆个参数是参数是否存在的bool值，可以⽤来判断参数是否存在。
		id,ok:=c.GetPostForm("id")
		if !ok {
			//参数不存在
			c.String(404, "id参数不存在")
		}
		// 获取PostForm的数组值
		names:=c.PostFormArray("names")
		c.JSON(200,gin.H{
			"name": name,
			"names": names,
			"id":id,
		})

	})
*/
    // ShouldBind函数 和go strust 进行绑定
	r.POST("/user", func(c *gin.Context) {
		//初始化User
		user := User{}
		//通过ShouldBind函数，将请求参数绑定到struct对象，处理json请求代码是⼀样的。
		//如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
		if c.ShouldBind(&user) == nil {
			//绑定成功，打印请求参数
			log.Println("user:", user)
		}
		//http请求返回⼀个字符串
		c.JSON(200, gin.H{
			"msg":             "parser Success",
			"请求的content-type": c.ContentType(),
			"解构后的数据":          user,
		})
	})

	// 启动服务
	r.Run()
}
