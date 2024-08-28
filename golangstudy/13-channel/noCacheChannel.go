package main

// import (
// 	"fmt"
// 	"time"
// )
//
// func main() {
// 	// 有缓存的channel
// 	ch := make(chan int, 3)
// 	go func() {
// 		defer fmt.Println("子 goroutine exit")
// 		for i := 0; i < 3; i++ {
// 			ch <- i
// 			fmt.Println("子 goroutine running i = ", i, "len = ", len(ch), " cap = ", cap(ch))
// 		}
// 	}()
//
// 	// channel 满了之后，再往里面写数据，会阻塞等待，channel为空，再从里面读取数据，也会阻塞等待
// 	time.Sleep(time.Second * 2)
// 	// 阻塞等待子 goroutine 结束
// 	for i := 0; i < 3; i++ {
// 		// 从ch中接收数据，并且赋值给num
// 		num := <-ch
// 		fmt.Println("num = ", num)
// 	}
// 	fmt.Println("main 结束")
// }
