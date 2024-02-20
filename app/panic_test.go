package main_test

import "testing"

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered in testPanic", r)
		}
	}()
	panic("TestPanic")
}
