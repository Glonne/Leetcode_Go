// 编写一个Golang程序，使用goroutine和channel实现两个并发任务之间的通信。
package main

import (
	"fmt"
	"time"
)

func task1(ch chan<- string) {
	fmt.Printf("task1 is running\n")
	ch <- "task1 is finished"
}

func task2(ch chan string) {
	fmt.Printf("task2 is running\n")
	time.Sleep(time.Second * 2)
	str := <-ch
	fmt.Printf("ch : %s", str)
	ch <- "task2 is finished"
}

func main() {
	ch := make(chan string, 0)
	go task1(ch)
	go task2(ch)
	time.Sleep(time.Second * 5)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
