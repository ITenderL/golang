package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// close可以关闭一个channel
		close(ch)
	}()
	for {
		// 判断channel是否关闭,ok为ture表示没有关闭，false表示channel关闭
		if data, ok := <-ch; !ok {
			break
		} else {
			println("data:", data)
		}
	}
	fmt.Println("Main Finished")
}
