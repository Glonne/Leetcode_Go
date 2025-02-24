package main

// 实现一个管道（pipeline），将多个channel串联起来，每个阶段对数据进行处理。
import (
	"fmt"
	"sync"
)

func generateData(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func doubleData(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func addTen(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num + 10
	}
	close(out)
}

func printResult(in <-chan int) {
	for num := range in {
		fmt.Println("Final result:", num)
	}
}

func main() {
	stage1 := make(chan int)
	stage2 := make(chan int)
	stage3 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		generateData(stage1)
	}()

	go func() {
		defer wg.Done()
		doubleData(stage1, stage2)
	}()
	go func() {
		defer wg.Done()
		addTen(stage2, stage3)
	}()
	go func() {
		defer wg.Done()
		printResult(stage3)
	}()
	wg.Wait()
}

// 分布式锁相关题目
// 使用Redis实现一个简单的分布式锁。
// 编写一个程序，使用etcd实现分布式锁。
// 实现一个基于Zookeeper的分布式锁。
// 使用Mutex和RWMutex实现一个本地锁，并比较它们的性能。
// 编写一个程序，模拟多个客户端竞争分布式锁的场景。
// 高级题目
// 使用Golang实现一个简单的RPC服务。
// 编写一个程序，使用Golang的http包实现一个简单的Web服务器。
// 使用Golang的sync.Pool来优化内存分配。
// 实现一个简单的负载均衡器，使用Golang的channel和goroutine。
// 使用Golang的pprof工具分析程序的性能瓶颈。
