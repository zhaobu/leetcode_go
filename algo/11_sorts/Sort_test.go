package _1_sorts

import (
	"fmt"
	"testing"
)

var (
	arr = []int{4, 5, 6, 3, 2, 1}
)

func TestBubbleSort(t *testing.T) {
	fmt.Println("排序前：", arr)
	BubbleSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestInsertionSort(t *testing.T) {
	fmt.Println("排序前：", arr)
	InsertionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestSelectionSort(t *testing.T) {
	fmt.Println("排序前：", arr)
	SelectionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}
