// 使用select语句处理多个channel的并发操作。
package main

import (
	"fmt"
	"time"
)

func producer1(ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Producer1: %d", i)
		time.Sleep(1 * time.Second)
	}
}

func producer2(ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Producer2: %d", i)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer1(ch1)
	go producer2(ch2)

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}

	fmt.Println("All messages received.")
}
