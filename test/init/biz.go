package main

import "fmt"

func init() {
	fmt.Println("biz init() :", a)
	a++
}

func add(a, b int) int {
	return a + b
}
