package main

import "fmt"

// interface是万能类型
func MyFunc(arg interface{}) {
	fmt.Println("MyFunc is Called")
	fmt.Println(arg)

	// 给Interface提供的断言机制
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	} else {
		fmt.Println("arg is not string type")
	}
}

type Book struct {
	author string
}

func main() {
	book := Book{"zhangsan"}
	MyFunc(book)

	MyFunc(123)
	MyFunc("asdf")

	MyFunc(false)
}
