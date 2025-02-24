package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	turn := 0
	var wg sync.WaitGroup

	numbers := []rune{'1', '2', '3', '4'}
	letters := []rune{'q', 'w', 'e', 'r'}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, num := range numbers {
			mu.Lock()
			for turn != 0 {
				mu.Unlock()
				mu.Lock()
			}
			fmt.Print(string(num))
			turn = 1
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for _, letter := range letters {
			mu.Lock()
			for turn != 1 {
				mu.Unlock()
				mu.Lock() // 手动等待另一协程释放锁
			}
			fmt.Print(string(letter))
			turn = 0 // 改变执行顺序为数字协程
			mu.Unlock()
		}
	}()

	wg.Wait() // 等待两个协程完成
}
