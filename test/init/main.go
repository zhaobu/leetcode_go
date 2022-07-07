package main

import "fmt"

func init() {
	fmt.Println("main init() 2:", a)
	a++
}

var a = 10

const b = 100

func main() {
	fmt.Printf("%d+%d=%d\n", a, b, add(a, b))
}

func init() {
	fmt.Println("main init() 1:", a)
	a++
}
