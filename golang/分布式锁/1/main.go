package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

var rdb *redis.Client
var ctx = context.Background()

const lockKey = "my_distributed_lock"
const lockTimeout = 10 * time.Second        // 锁的过期时间，防止死锁
const retryTimeout = 100 * time.Millisecond // 锁获取失败时重试的间隔
const maxRetries = 5                        // 锁获取失败时最大重试次数

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 如果有密码，设置密码
		DB:       0,                // 使用的数据库
	})
}

func acquireLock(lockID string) (bool, error) {

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

func main() {

	lockID := uuid.New().String()

	var acquired bool
	var err error

	// 重试获取锁
	for retries := 0; retries < maxRetries; retries++ {
		acquired, err = acquireLock(lockID)
		if err != nil {
			log.Fatalf("Error acquiring lock: %v", err)
		}
		if acquired {
			fmt.Println("Lock acquired, performing task...")
			break
		}
		// 锁已被占用，稍后重试
		fmt.Println("Lock is already acquired, retrying...")
		time.Sleep(retryTimeout)
	}

	if !acquired {
		log.Fatalf("Could not acquire lock after %d retries", maxRetries)
		return
	}
	time.Sleep(5 * time.Second)
	err = releaseLock(lockID)
	if err != nil {
		log.Fatalf("Error releasing lock: %v", err)
	} else {
		fmt.Println("Lock released.")
	}
}
