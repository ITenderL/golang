package main

import (
	"fmt"
	"time"
)

func main() {
	// go func() {
	// 	defer fmt.Println("A defer")
	// 	func() {
	// 		defer fmt.Println("B defer")
	// 		// 退出当前 goroutine
	// 		runtime.Goexit()
	// 		fmt.Println("B")
	// 	}()
	// 	fmt.Println("A")
	// }()

	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
