package main_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestAsyncService(t *testing.T) {

	wg := new(sync.WaitGroup)

	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start goroutine1")
		ch <- "Hello"
		fmt.Println("goroutine1 done")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start goroutine2")
		recieved := <-ch
		fmt.Println("goroutine2 recieved:", recieved)
		fmt.Println("goroutine2 done")
	}()

	wg.Wait()
}
