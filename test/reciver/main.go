package main

import (
	"fmt"
	"runtime/debug"
)

type MySlice []int

func (m MySlice) append1(i int) {
	fmt.Printf(string(debug.Stack()))
	m = append(m, i)
}

func (m *MySlice) append2(i int) {
	fmt.Printf(string(debug.Stack()))
	*m = append(*m, i)
}
func main() {
	m := make(MySlice, 2, 4)
	fmt.Printf("m=%+v\n", m)
	m.append1(1)
	m.append2(2)
	fmt.Printf("m=%+v\n", m)
}
