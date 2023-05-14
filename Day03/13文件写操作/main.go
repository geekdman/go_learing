package main

import (
	"fmt"
	"os"
)

/*
os.OpenFile函数

*/

func WriteBytesOrstr(file *os.File) {
	str := "满江红\n"
	//file.Write([]byte(str))
	file.WriteString(str)
}
func WriteFile() {
	data := `
	满江红
	怒发冲冠，凭栏处，潇潇雨歇
`
	os.WriteFile("满江红3.txt", []byte(data), 0666)

}

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	WriteBytesOrstr(file)
	WriteFile()
}
