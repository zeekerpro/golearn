package main_test

import (
	"fmt"
	"testing"
)

type Talker interface {
	Talk() string
}

type Person struct{ Name string }

func (p Person) Talk() string {
	p.Name = "wht"
	return "Hi, my name is " + p.Name
}

type Robot struct{ Model string }

func (r *Robot) Talk() string {
	r.Model = "god"
	return "I am model " + r.Model
}

func Greet(t Talker) {
	fmt.Println(t.Talk())
}

func TestGreet(t *testing.T) {
	p := Person{Name: "Alice"}
	Greet(p)
	Greet(&p)
	t.Log(p.Name)

	r := Robot{Model: "T-800"}
	Greet(&r)
	t.Log(r.Model)

}
