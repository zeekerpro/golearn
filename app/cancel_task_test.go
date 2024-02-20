package main_test

import (
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
