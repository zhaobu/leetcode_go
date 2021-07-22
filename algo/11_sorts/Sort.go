package _1_sorts

import "fmt"

/*
冒泡排序、插入排序、选择排序
*/

//冒泡排序，a是数组，n表示数组大小
func BubbleSort(a []int, n int) {
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ { //循环n次
		swap := false                //标记是否有交换
		for j := 0; j < n-i-1; j++ { //每轮比上一轮少比较一次
			if a[j+1] < a[j] {
				swap = true
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
		fmt.Printf("a=%+v\n", a)
		if !swap {
			return
		}
	}
}

// 插入排序，a表示数组，n表示数组大小
func InsertionSort(a []int, n int) {
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ { //依次遍历未排序区间
		for j := i; j > 0; j-- { //从已排序区间最后一个位置开始查找
			if a[j] >= a[j-1] { //直到找到前一个元素比自己小时才停下
				break
			}
			a[j], a[j-1] = a[j-1], a[j] //如果前一个元素比自己大,就把前一个元素往后移动一个位置
		}
		fmt.Printf("a=%+v\n", a)
	}
}

// 选择排序，a表示数组，n表示数组大小
func SelectionSort(a []int, n int) {
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i //最小值所在索引
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
		fmt.Printf("a=%+v\n", a)
	}
}
