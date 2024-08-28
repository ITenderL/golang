package main

// import (
// 	"fmt"
// 	"time"
// )
//
// // 子goroutine
// func newTask() {
// 	i := 0
// 	for {
// 		i++
// 		fmt.Println("new goroutine: i = ", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }
//
// func main() {
// 	// 创建新任务go程
// 	go newTask()
// 	i := 0
// 	for {
// 		i++
// 		fmt.Println("main goroutine: i = ", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }
