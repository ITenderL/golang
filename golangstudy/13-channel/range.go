package main

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()
// 	// 遍历channel, 从channel中读取数据，直到channel关闭
// 	for data := range ch {
// 		println("data:", data)
// 	}
// }
