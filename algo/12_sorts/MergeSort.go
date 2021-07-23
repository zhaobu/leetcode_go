package _2_sorts

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)>>1
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, 0, end-start+1)
	i, j := start, mid+1
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp = append(tmp, arr[i])
	}
	for ; j <= end; j++ {
		tmp = append(tmp, arr[j])
	}

	copy(arr[start:end+1], tmp)
}
