package main

import "fmt"

func printMyArray(arr []int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
	arr[0] = 100
}

// 动态数组, 切片传递的时数组的引用
// func main() {
// 	myArray := []int{5, 6, 7, 8, 9}
// 	fmt.Printf("myArray type = %T\n", myArray)
// 	printMyArray(myArray)
//
// 	for i := range myArray {
// 		fmt.Println(myArray[i])
// 	}
// 	// 声明一个slice，并且初始化，初始值为1，2，3，长度为3
// 	// slice := []int{1, 2, 3}
// 	// 声明slice是一个切片但没有分配空间
// 	var slice1 []int
// 	// 开辟3个空间，默认值为0
// 	slice1 = make([]int, 3)
// 	// 声明slice是一个切片，长度为3，初始值0
// 	// var slice2 = make([]int, 3)
//
// 	if slice1 == nil {
// 		fmt.Println("slice1 是一个空切片")
// 	} else {
// 		fmt.Println("slice1是有空间的")
// 	}
// }
