package main

import "fmt"

func main() {
	m := 10
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		arr[i] = i + 1
	}
	fmt.Printf("arr=%+v\n", arr)
}
