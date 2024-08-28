package main

import "fmt"

// const来定义枚举,iota只能配合const()来使用
const (
	// BEIJING 可以在const的添加一个iota，每行的iota都会累加，第一行的iota默认值为0
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

func main() {
	// 常量(只读属性)，不允许修改
	const name = "asdf"
	const length = 100

	fmt.Print("BEIJING = ", BEIJING)
	fmt.Print("SHANGHAI = ", SHANGHAI)
	fmt.Print("BEIJING = ", SHENZHEN)
}
