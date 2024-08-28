package main

import (
	"fmt"
	"reflect"
)

// import (
// 	"fmt"
// 	"reflect"
// )
//
// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }
//
// func (user User) Call() {
// 	fmt.Println("user is called....")
// 	fmt.Printf("%v\n", user)
// }
//
// func main() {
// 	user := User{1, "haha", 20}
// 	DoFieldMethod(user)
// }

func DoFieldMethod(input interface{}) {
	// 获取type类型
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is ", inputType.Name())
	// 获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is ", inputValue)

	// 通过type获取里面的字段
	// 通过reflect获取type，通过type获取NumField，遍历
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s : %v\n", method.Name, method.Type)
	}
}
