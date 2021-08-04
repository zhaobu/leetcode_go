package LinearSort

import (
	"fmt"
	Sort "leetcode/algo/12_sorts"
)

// 桶排序

// 获取待排序数组中的最大值
func getMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func getMaxMin(a []int) (int, int) {
	max := a[0]
	min := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	return max, min
}

func BucketSort(a []int) {
	num := len(a)
	if num <= 1 {
		return
	}
	max, min := getMaxMin(a)      //找到数组中最大,最小值
	buckets := make([][]int, num) // 二维切片,创建和a长度一样的个数的桶
	diff := max - min

	index := 0
	for i := 0; i < num; i++ {
		index = (a[i] - min) * (num - 1) / diff       // 桶序号
		buckets[index] = append(buckets[index], a[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			Sort.QuickSort(buckets[i]) // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}

}

// 桶排序简单实现
func BucketSortSimple(source []int) {
	if len(source) <= 1 {
		return
	}
	array := make([]int, getMax(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]]++
	}
	fmt.Println(array)
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(source, c)

}
