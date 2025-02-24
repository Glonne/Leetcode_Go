package main

import (
	"fmt"
	"sync"
)

func fanOut(in <-chan int, outs []chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range in {
		for _, out := range outs {
			go func(out chan int) {
				out <- data
			}(out)
		}
	}
}

func main() {

	in := make(chan int)

	out1 := make(chan int)
	out2 := make(chan int)
	out3 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go fanOut(in, []chan int{out1, out2, out3}, &wg)

	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		for data := range out1 {
			fmt.Printf("out1: %d\n", data)
		}
	}()

	go func() {
		for data := range out2 {
			fmt.Printf("out2: %d\n", data)
		}
	}()

	go func() {
		for data := range out3 {
			fmt.Printf("out3: %d\n", data)
		}
	}()

	wg.Wait()

	close(out1)
	close(out2)
	close(out3)
}
