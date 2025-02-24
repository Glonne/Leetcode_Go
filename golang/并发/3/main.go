package main

import (
	"fmt"
	"sync"
	"time"
)

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Printf("消费了 %d\n", item)
		time.Sleep(time.Second * 2)
	}
}

func producer(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(time.Second)
	}
	close(ch)
	defer wg.Done()
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("All tasks are finished.")
}
