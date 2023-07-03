package main

import (
	"fmt"
	"time"
)

/*

时间戳：从１９９７／０１／０１／０：０：０　开始计算

*/

func bar(x,y int)   {
	fmt.Println(x+y)
}
func timer(f func(x,y int)) {
	start := time.Now().Unix()
	f(x,y)
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func main() {
	a,b := 10,20
	timer(bar(a,b))
}