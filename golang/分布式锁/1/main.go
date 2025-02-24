package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

var rdb *redis.Client
var ctx = context.Background()

const lockKey = "distributed_task_lock"
const lockTimeout = 10 * time.Second        // 锁的过期时间
const retryTimeout = 100 * time.Millisecond // 锁获取失败时重试的间隔
const maxRetries = 5                        // 锁获取失败时最大重试次数

// 初始化Redis客户端
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 如果有密码，设置密码
		DB:       0,                // 使用的数据库
	})
}

// 获取分布式锁
func acquireLock(lockID string) (bool, error) {
	// 尝试获取锁，如果锁不存在，则设置它
	success, err := rdb.SetNX(ctx, lockKey, lockID, lockTimeout).Result()
	if err != nil {
		return false, err
	}
	return success, nil
}

// 检查当前持有者是否为当前进程
func isLockOwner(lockID string) (bool, error) {
	currentOwner, err := rdb.Get(ctx, lockKey).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return currentOwner == lockID, nil
}

// 释放锁，确保只有锁的持有者能够释放锁
func releaseLock(lockID string) error {
	// 确认当前持有锁的进程是我们自己的 lockID
	isOwner, err := isLockOwner(lockID)
	if err != nil {
		return err
	}
	if !isOwner {
		return fmt.Errorf("current process is not the lock owner")
	}

	// 只有持有锁的进程才能删除锁
	_, err = rdb.Del(ctx, lockKey).Result()
	return err
}

// 模拟执行任务
func executeTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 生成一个唯一的锁ID，表示当前进程
	lockID := uuid.New().String()

	var acquired bool
	var err error

	// 重试获取锁
	for retries := 0; retries < maxRetries; retries++ {
		acquired, err = acquireLock(lockID)
		if err != nil {
			log.Printf("Error acquiring lock for task %d: %v", taskID, err)
			return
		}
		if acquired {
			fmt.Printf("Task %d acquired the lock, performing task...\n", taskID)
			break
		}
		// 锁已被占用，稍后重试
		fmt.Printf("Task %d: Lock is already acquired, retrying...\n", taskID)
		time.Sleep(retryTimeout)
	}

	if !acquired {
		log.Printf("Task %d could not acquire the lock after %d retries\n", taskID, maxRetries)
		return
	}

	// 模拟任务执行
	time.Sleep(5 * time.Second)

	// 释放锁
	err = releaseLock(lockID)
	if err != nil {
		log.Printf("Error releasing lock for task %d: %v", taskID, err)
	} else {
		fmt.Printf("Task %d released the lock.\n", taskID)
	}
}

func main() {
	var wg sync.WaitGroup

	// 启动多个并发的任务
	for i := 1; i <= 3; i++ { // 启动 3 个任务
		wg.Add(1)
		go executeTask(i, &wg)
	}

	// 等待所有任务完成
	wg.Wait()
}
