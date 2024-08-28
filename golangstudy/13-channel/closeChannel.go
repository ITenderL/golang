package main

// import "fmt"
//
// func main() {
// 	// 创建channel,channel需要用make来初始化
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 		// close可以关闭一个channel，关闭后channel不能再写入数据,
// 		// 如果关闭一个已经关闭的channel，会引发panic
// 		// 关闭channel后，如果channel中还有数据，可以读取到数据
// 		close(ch)
// 	}()
// 	for {
// 		// 判断channel是否关闭,ok为ture表示没有关闭，false表示channel关闭
// 		if data, ok := <-ch; !ok {
// 			break
// 		} else {
// 			println("data:", data)
// 		}
// 	}
// 	fmt.Println("Main Finished")
// }
