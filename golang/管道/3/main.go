// 实现一个fan-in模式，将多个channel的数据合并到一个channel中。
package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < 3; i++ {
			out <- fmt.Sprintf("Producer %d: Message %d", id, i)
			time.Sleep(time.Duration(id) * 200 * time.Millisecond)
		}
	}()
	return out
}

func fanIn(inputs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				out <- msg
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	ch1 := producer(1)
	ch2 := producer(2)
	ch3 := producer(3)

	merged := fanIn(ch1, ch2, ch3)

	for msg := range merged {
		fmt.Println(msg)
	}

	fmt.Println("All data processed.")
}
