package _2_sorts

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	fmt.Printf("排序前arr=%+v\n", arr)
	MergeSort(arr)
	fmt.Printf("排序后arr=%+v\n", arr)

	arr = []int{5, 4, 3, 2, 1}
	fmt.Printf("排序前arr=%+v\n", arr)
	MergeSort(arr)
	fmt.Printf("排序后arr=%+v\n", arr)

	arr = createRandomArr(100)
	fmt.Printf("排序前arr=%+v\n", arr)
	MergeSort(arr)
	fmt.Printf("排序后arr=%+v\n", arr)
}
