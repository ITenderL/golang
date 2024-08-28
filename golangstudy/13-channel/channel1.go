package main

// import "fmt"
//
// func main() {
// 	c := make(chan int)
// 	go func() {
// 		defer fmt.Println("goroutine 结束")
// 		fmt.Println("goroutine 开始")
// 		// 发送数据给c
// 		c <- 1
// 	}()
// 	// 接收数据c，赋值给num
// 	num := <-c
// 	fmt.Println("num = ", num)
// }
