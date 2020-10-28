package Solution

/*
[冒泡排序](https://www.geekxh.com/2.0.%E6%8E%92%E5%BA%8F%E7%B3%BB%E5%88%97/1.bubbleSort.html)
*/
func BubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		sort := true
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				sort = false //发生了交换说明还不是有序
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if sort {
			break
		}
	}
	return nums
}

/*
[选择排序](https://www.geekxh.com/2.0.%E6%8E%92%E5%BA%8F%E7%B3%BB%E5%88%97/2.selectionSort.html#_1-%E7%AE%97%E6%B3%95%E6%AD%A5%E9%AA%A4)
*/
func SelectionSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		m := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[m] {
				m = j
			}
		}
		nums[i], nums[m] = nums[m], nums[i]
	}
	return nums
}

/*
[插入排序](https://www.geekxh.com/2.0.%E6%8E%92%E5%BA%8F%E7%B3%BB%E5%88%97/3.insertionSort.html#_1-%E7%AE%97%E6%B3%95%E6%AD%A5%E9%AA%A4)
*/
func InsertionSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
	return nums
}

/*
[希尔排序](https://blog.csdn.net/qq_39207948/article/details/80006224)
*/
func ShellSort(nums []int) []int {
	n := len(nums)

	//进行分组,最开始时的增量gap为数组长度的一半
	for gap := n / 2; gap > 0; gap /= 2 {
		/*
			对各个分组进行插入排序的时候并不是先对一个组进行排序完再对另一个组进行排序,而是轮流对每个组进行插入排序
		*/
		for i := gap; i < n; i++ {
			//i:代表即将插入的元素下标
			// j代表和i同组的每个左边的元素,把nums[i]插入到该组正确的位置
			for j := i; j >= gap && nums[j] < nums[j-gap]; j -= gap {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
			}
		}
	}
	return nums
}

/*
[归并排序](https://www.geekxh.com/2.0.%E6%8E%92%E5%BA%8F%E7%B3%BB%E5%88%97/5.mergeSort.html#_3-%E5%8A%A8%E5%9B%BE%E6%BC%94%E7%A4%BA)
*/
func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	middle := n / 2
	left, right := nums[0:middle], nums[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	n1, n2 := len(left), len(right)
	res := make([]int, 0, n1+n2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	if i < n1 {
		res = append(res, left[i:]...)
	} else if j < n2 {
		res = append(res, right[j:]...)
	}
	return res
}

/*
[快速排序](https://www.geekxh.com/2.0.%E6%8E%92%E5%BA%8F%E7%B3%BB%E5%88%97/6.quickSort.html#_1-%E7%AE%97%E6%B3%95%E6%AD%A5%E9%AA%A4)
*/
func QuickSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	quickSort(nums, 0, n-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		i, j, key := start, end, nums[start]
		for i < j {
			for i < j && key < nums[j] {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			for i < j && nums[i] < key {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				j--
			}
		}
		quickSort(nums, start, i-1)
		quickSort(nums, i+1, end)
	}
}

/*
[堆排序](https://www.geekxh.com/2.0.%E6%8E%92%E5%BA%8F%E7%B3%BB%E5%88%97/6.quickSort.html#_1-%E7%AE%97%E6%B3%95%E6%AD%A5%E9%AA%A4)
可以结合https://www.cnblogs.com/ssyfj/p/9512451.html理解
*/
func HeapSort(nums []int) []int {

	/*
		low表示二叉树中最后一个非叶结点下标,high表示数组最后一个元素的下标
		下标从0开始的数组构建成的二叉树,父结点为i,则左孩子为2*i+1,右孩子为2*i+2
	*/
	heapAdjust := func(arr []int, low, high int) {
		i, tmp := low, arr[low]
		for j := 2*i + 1; j <= high; j = 2*j + 1 { //逐渐去找左右孩子结点
			//找到两孩子结点中最大的
			if j < high && arr[j] < arr[j+1] {
				j++
			}
			//父节点和孩子最大的进行判断，调整，变为最大堆
			if tmp >= arr[j] {
				break
			}
			//将父节点数据变为最大的，将原来的数据还是放在temp中，
			arr[i] = arr[j]
			i = j
		}
		arr[i] = tmp
	}

	n := len(nums)
	/*
		首先将无序数列转换为大顶堆,采取下标从0开始的数组构建:
		注意由于是完全二叉树，从最后一个非叶结点(n/2-1)开始构建
	*/
	for i := n/2 - 1; i >= 0; i-- {
		heapAdjust(nums, i, n-1)
	}
	/*
		上面大顶堆已经构造完成，我们现在需要排序，每次将二叉树根结点也就是最大的元素放入最后
		然后将剩余元素重新构造大顶堆，将最大元素放在剩余最后
	*/
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapAdjust(nums, 0, i-1)
	}
	return nums
}
