package _5_binarysearch

//普通二分查找
func BinarySearch1(a []int, v int) int {
	if len(a) < 1 {
		return -1
	}
	low, high := 0, len(a)-1
	for low < high { //注意是 low<=high，而不是 low
		mid := low + (high-low)>>1 //low + (high-low)>>1  可以防止(low+high)/2这种写法溢出
		if v < a[mid] {
			high = mid - 1
		} else if v > a[mid] {
			low = mid + 1
		} else if a[mid] == v {
			return mid
		}
	}
	if a[low] == v {
		return low
	}
	return -1
}

//普通二分查找递归实现
func BinarySearchRecursive1(a []int, v int) int {
	if len(a) < 1 {
		return -1
	}

	return bs(a, v, 0, len(a)-1)
}
func bs1(a []int, v, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if a[mid] < v {
		return bs1(a, v, mid+1, high)
	} else if a[mid] > v {
		return bs1(a, v, low, mid-1)
	} else {
		return mid
	}
}

//查找第一个等于给定值的元素
func BinarySearchFirst1(a []int, v int) int {
	if len(a) < 1 {
		return -1
	}
	low, high := 0, len(a)-1
	for low < high {
		mid := low + (high-low)>>1
		if a[mid] < v {
			low = mid + 1
		} else {
			//继续往左找
			high = mid - 1
		}
	}
	if a[low] == v {
		return low
	}
	return -1
}

//查找最后一个值等于给定值的元素
func BinarySearchLast1(a []int, v int) int {
	if len(a) < 1 {
		return -1
	}

	low, high := 0, len(a)-1
	for low < high {
		mid := low + (high-low)>>1
		if a[mid] < v {
			low = mid + 1
		} else if a[mid] > v {
			high = mid - 1
		} else {
			//继续往右查找
			low = mid
		}
	}
	if a[low] == v {
		return low
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchFirstGT1(a []int, v int) int {
	if len(a) < 1 {
		return -1
	}
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v { //当a[mid]>=v时,判断mid左边的值是否<v
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素
func BinarySearchLastLT1(a []int, v int) int {
	if len(a) < 1 {
		return -1
	}
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] <= v {
			if mid == len(a)-1 || a[mid+1] > v { //当a[mid]<=v时,判断mid右边的值是否>v
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] { //[0,mid]是升序的
			if target >= nums[0] && target < nums[mid] { //target在[0,mid)
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { //[mid,len(nums)-1]是升序的
			if target <= nums[len(nums)-1] && target > nums[mid] { //target在(mid,len(nums)-1]
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
