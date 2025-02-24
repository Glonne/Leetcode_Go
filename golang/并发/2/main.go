// 使用sync.WaitGroup来等待多个goroutine完成
package main

import (
	"fmt"
	"sync"
)

func task(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("task %d is running\n", num)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go task(&wg, i)
	}
	wg.Wait()
}
