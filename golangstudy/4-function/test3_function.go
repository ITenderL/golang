package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Print("a = ", a)
	fmt.Print(", b = ", b, "\n")
	c := 100
	return c
}

// 返回多个返回值，匿名返回
func foo2(a string, b int) (int, int) {
	fmt.Print("a = ", a)
	fmt.Print(", b = ", b, "\n")
	return 66, 77
}

// 返回多个返回值，有名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Print("a = ", a)
	fmt.Print(", b = ", b, "\n")
	r1 = 1000
	r2 = 2000
	return
}

// 返回多个返回值，有名称的
func foo4(a string, b int) (r3, r4 int) {
	fmt.Print("a = ", a)
	fmt.Print(", b = ", b, "\n")
	r3 = 3000
	r4 = 4000
	return
}

func main() {
	res := foo1("haha", 555)
	fmt.Print("res = ", res, "\n")
	res1, res2 := foo2("lala", 888)
	fmt.Print("res1 = ", res1, " , res2 = ", res2)
	r1, r2 := foo3("jjjj", 999)
	fmt.Print("r1 = ", r1, " , r2 = ", r2)
	r3, r4 := foo3("jjjj", 999)
	fmt.Print("r3 = ", r3, " , r4 = ", r4)
}
