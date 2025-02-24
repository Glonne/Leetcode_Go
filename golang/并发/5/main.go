// 编写一个程序，使用context包来取消一个长时间运行的goroutine
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task was cancelled.")
			return
		default:
			fmt.Println("Task is running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go longRunningTask(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

// 用于管理 Goroutine 的生命周期。它可以帮助我们实现超时控制、取消任务
// 以及在 Goroutine 之间传递请求范围的数据。context 主要用于防止 Goroutine 泄漏，确保在任务完成或不需要时能够正确终止 Goroutine，提高资源利用率。
