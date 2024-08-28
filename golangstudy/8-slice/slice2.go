package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 追加一个元素，长度len = 4, cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// cap已经满了，超出cap，则cap自动扩容为原来的2倍
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	s := []int{1, 2, 3}
	// 下表[)，左闭右开
	s1 := s[0:2]
	fmt.Println(s1)
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	// 深拷贝，重新开辟内存空间，指向新的地址
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
}
