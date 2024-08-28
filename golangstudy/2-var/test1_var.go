package main

import "fmt"

// 声明全局变量，方法1、2、3没有区别
var aa int
var bb int = 100

var dd = 100

// 用方法四声明全局变量,:=z只能在函数体内使用

/*
四种变量声明方式
*/
func main() {
	// 方法一：声明一个变量，没有初始值，默认为0
	var a int
	fmt.Print("a = ", a)

	// 方法二：给定默认值
	var b int = 100
	fmt.Print("b = ", b)

	// 方法三：初始化的时候可省去类型
	var c = 100
	fmt.Print("c = ", c)
	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	// 方法四：生气var关键字
	e := 100
	fmt.Printf("e = %d, type of e = %T\n", e, e)

	f := "asdf"
	fmt.Print("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14159
	fmt.Printf("g = %f, type of g = %T\n", g, g)
	// 全局变量
	fmt.Print(aa, bb, dd)
	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Print("xx = ", xx, "yy = ", yy)

	var jj, kk = 100, "asdf"
	fmt.Print("jj = ", jj, "kk = ", kk)

	// 多行多变量
	var (
		gg = 100
		hh = "asdf"
	)
	fmt.Print("gg = ", gg, ", hh = ", hh)
}
