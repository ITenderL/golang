package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func deferFunc() int {
	fmt.Println("deferFunc,,,,,,,,,")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc,,,,,,,,,")
	return 0
}

func deferAndReturnFunc() int {
	defer deferFunc()
	return returnFunc()
}

// defer为后执行的方法
// defer执行为压栈的形式，FILO
func main() {
	// 7-defer func1()
	// 7-defer func2()
	// 7-defer func3()
	// fmt.Println("Hello GO!")
	// return先执行，defer后执行
	deferAndReturnFunc()
}
