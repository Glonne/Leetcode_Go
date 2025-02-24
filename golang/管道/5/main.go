// 使用channel实现一个简单的任务调度器
package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID      int
	Message string
}

func worker(id int, tasks <-chan Task, done chan<- bool) {
	for task := range tasks {
		fmt.Printf("Worker %d: Processing task %d - %s\n", id, task.ID, task.Message)
		time.Sleep(time.Second) // 模拟任务执行的时间
	}
	done <- true
}

func main() {

	tasks := make(chan Task, 10)
	done := make(chan bool, 3)
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, done)
	}
	for i := 1; i <= 10; i++ {
		tasks <- Task{ID: i, Message: fmt.Sprintf("Task %d", i)}
	}
	close(tasks)
	for i := 1; i <= 3; i++ {
		<-done
	}
	fmt.Println("All tasks are completed.")
}
