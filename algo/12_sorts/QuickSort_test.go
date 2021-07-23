package _2_sorts

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func createRandomArr(length int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 4}
	fmt.Printf("排序前arr=%+v\n", arr)
	QuickSort(arr)
	fmt.Printf("排序后arr=%+v\n", arr)

	arr = createRandomArr(10)
	fmt.Printf("排序前arr=%+v\n", arr)
	QuickSort(arr)
	fmt.Printf("排序后arr=%+v\n", arr)
}
