package main

import "fmt"

func printMap(cityMap map[string]string) {
	// 是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func main() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap是一个空map")
	}
	// 第一种声明方式，空的容量，需要make声明容量
	myMap = make(map[string]string, 10)
	myMap["one"] = "go"
	myMap["two"] = "Java"
	fmt.Println(myMap)
	// 第二种方式
	myMap2 := make(map[int]string)
	myMap2[1] = "go"
	myMap2[2] = "Java"
	fmt.Println(myMap2)
	// 第三种方式：初始话的map
	myMap3 := map[string]string{
		"one":   "Java",
		"two":   "GO",
		"three": "C++",
	}
	fmt.Println(myMap3)

	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
	// 修改
	cityMap["USA"] = "HSD"
	// 删除
	delete(cityMap, "China")
	printMap(cityMap)

	// for key, value := range cityMap {
	// 	fmt.Println("key = ", key)
	// 	fmt.Println("value = ", value)
	// }
}
