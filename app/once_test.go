package main_test

import (
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singletonInstance *Singleton
var once sync.Once

func getSingletonInstance() *Singleton {
	once.Do(func() {
		println("Creating singleton instance")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

// go test -v -run=TestGetSingletonInstance
func TestGetSingletonInstance(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			instance := getSingletonInstance()
			t.Logf("%p", unsafe.Pointer(instance))
			wg.Done()
		}()
	}
	wg.Wait()
}
