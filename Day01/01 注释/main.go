package main //定义包名，其中main 是一个模块的入口

/*
	多行注释
*/

/*
  使用 import 字段导入包,格式为：import + <包名>
*/
import "fmt"

// 定义主函数main ，格式为： func + <包名>
func main() {
	// 使用"//"进行单行注释，注释是不会
	/*
	   这是一个多行注释
	   这是一个多行注释
	   这是一个多行注释
	*/
	fmt.Println("Hello World")
}
