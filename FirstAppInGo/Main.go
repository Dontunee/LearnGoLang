package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	func() {
		fmt.Println("Hello Go")
	}()
}

func sayMessage(message string) {
	fmt.Println(message)
}
