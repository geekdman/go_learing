package main

import (
	"fmt"
	"reflect"
)

func test(customers *[]map[string]interface{}) {
	fmt.Println(reflect.TypeOf(customers))
	fmt.Println((*customers)[0]["name"])
	//for i, customerMap := range customers {
	//	if customerMap["name"] == name {
	//		fmt.Println(customers[i])
	//	}
	//}
}

func main() {

	var customers []map[string]interface{}
	customer := map[string]interface{}{"name": "tom", "age": 18, "height": 180}

	customers = append(customers, customer)
	fmt.Println(customers)
	test(&customers)
}
