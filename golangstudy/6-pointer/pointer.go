package main

import "fmt"

// 值传递
/*func swap(a int, b int) {
	2-var temp int
	temp = a
	a = b
	b = temp
}*/

// 地址传递
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a, b = 100, 200
	swap(&a, &b)
	fmt.Println("a = ", a, " b = ", b)

	// 二级指针
	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)

}
