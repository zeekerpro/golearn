package main_test

import (
	"sync"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(i)
		}()
		time.Sleep(1 * time.Second)
	}
}

func TestRoutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
		time.Sleep(1 * time.Second)
	}
}

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(1 + time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
		}()
	}

	wg.Wait()

	t.Logf("counter = %d", counter)
}
