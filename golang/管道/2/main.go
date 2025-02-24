package main

import (
	"fmt"
	"time"
)

// 执行任务的函数
func doTask(taskID int, ch chan struct{}) {
	defer func() {
		<-ch
	}()

	fmt.Printf("Task %d is starting...\n", taskID)
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d is done.\n", taskID)
}

func main() {

	concurrencyLimit := 3
	ch := make(chan struct{}, concurrencyLimit)

	for i := 1; i <= 10; i++ {
		ch <- struct{}{}
		go doTask(i, ch)
	}

	time.Sleep(12 * time.Second)
}
