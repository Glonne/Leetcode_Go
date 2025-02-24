package main

import (
	"fmt"
	"sync"
)

func main() {
	numberCh := make(chan bool, 1) // 控制数字输出的 channel
	letterCh := make(chan bool)    // 控制字母输出的 channel
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		numbers := []rune{'1', '2', '3', '4'}
		for _, n := range numbers {
			<-numberCh // 等待信号，才能执行
			fmt.Print(string(n))
			letterCh <- true // 发送信号，让字母协程执行
		}
	}()

	go func() {
		defer wg.Done()
		letters := []rune{'q', 'w', 'e', 'r'}
		for _, l := range letters {
			<-letterCh // 等待信号，才能执行
			fmt.Print(string(l))
			if l != 'r' {
				numberCh <- true // 发送信号，让数字协程执行
			}
		}
	}()
	numberCh <- true
	wg.Wait()     // 等待所有 goroutine 结束
	fmt.Println() // 换行
}
