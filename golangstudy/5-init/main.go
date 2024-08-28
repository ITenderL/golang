package main

import (
	// 匿名导报，无法使用当前包的方法，可以执行init方法
	_ "golangstudy/5-init/lib1"
	// 别名导报，使用别名调用
	// mylib2 "golangstudy/5-init/lib2"
	// 把包导入到当前包，不要轻易使用
	. "golangstudy/5-init/lib2"
)

// 匿名导包，别名
func main() {
	// lib1.TestLib1()
	// lib2.TestLib2()
	TestLib2()
}
