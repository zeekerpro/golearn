package main_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChannel chan struct{}) bool {
	select {
	case <-cancelChannel:
		return true
	default:
		return false
	}
}

func cancelTask(cancelChannel chan struct{}) {
	close(cancelChannel)
}

func TestCancelTask(t *testing.T) {
	cancelChannel := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelChannel chan struct{}) {
			for {
				if isCancelled(cancelChannel) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			t.Log(i, "Cancelled")
		}(i, cancelChannel)
	}
	cancelTask(cancelChannel)
	time.Sleep(time.Second * 1)
}

func operation(ctx context.Context, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Println("Operation finished")
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
	}
}

func TestContext(t *testing.T) {
	// 创建一个可取消的context
	ctx, cancel := context.WithCancel(context.Background())

	// 模拟一个需要一定时间完成的操作
	go operation(ctx, 3*time.Second)

	// 取消context
	time.Sleep(1 * time.Second) // 等待一段时间
	cancel()                    // 现在取消操作

	// 给操作足够的时间来响应取消
	time.Sleep(2 * time.Second)
}
