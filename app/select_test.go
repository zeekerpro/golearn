package main_test

import (
	"sync"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {

	wg := new(sync.WaitGroup)

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(1)
	go func() {
		ch1 <- 1
	}()

	wg.Add(1)
	go func() {
		ch2 <- 2
	}()

	select {
	case val, ok := <-ch1:
		if !ok {
			t.Log("ch1 closed")
		}
		t.Log(val)
	case val, ok := <-ch2:
		if !ok {
			t.Log("ch2 closed")
		}
		t.Log(val)
	case <-time.After(time.Second * 10):
		t.Error("Timeout")
	}

	close(ch1)
	close(ch2)

	wg.Wait()

}
