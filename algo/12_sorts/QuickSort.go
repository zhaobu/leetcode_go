package _2_sorts

// QuickSort is quicksort methods for golang
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end] //使用pivot把arr分为左右2部分

	i := start //arr[i]表示已处理区间的尾部，左边是小于pivot的，右边是大于pivot的

	for j := start; j <= end; j++ {
		if arr[j] < pivot { //遍历未处理区，把小于分区元素的都移动到已处理区后面
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
