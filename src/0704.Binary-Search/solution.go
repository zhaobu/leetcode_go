package Solution

/*[labuladong二分法讲解 ](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie)
使用左闭右闭区间l<=r,并且r=len(nums)-1
*/
func Search1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	/*
		结束条件left == right+1,l,r的取值范围是[0,len(nums)-1]
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 15}, expect: -1}求最大值,r=5,l=6
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 4}, expect: -1}中间值,r=2,l=3
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: -2}, expect: -1}求最小值,r=-1,l=0
		r:[-1,len(nums)-1]
		l:[0,len(nums)]
		运行到此处时的情况:
		1: l=len(nums),r=len(nums)-1并且nums[r]!=target,此时l和r取得最大值
		2: l=r+1并且l和r都是有效下标,但是nums[l]!=target,nums[r]!=target
		3:l=0,r=-1并且nums[0]!=target,此时l和r取得最小值
		综合三种情况都是return -1
	*/
	return -1
}

/*
[labuladong二分法讲解 ](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie)
使用左闭右开区间l<r,并且r=len(num)
*/
func Search2(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		}
	}
	/*
		结束条件left == right,l,r的取值范围是[0,len(nums)]
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 15}, expect: -1}求最大值,r=6,l=6
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 4}, expect: -1}中间值,r=3,l=3
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: -2}, expect: -1}求最小值,r=0,l=0
		r:[0,len(nums)]
		l:[0,len(nums)]
		运行到此处时的情况:
		1: l=len(nums),r=len(nums),此时l和r取得最大值
		2: l=r并且l和r都是有效下标,但是nums[l]!=target,nums[r]!=target
		3:l=0,r=0并且nums[0]!=target,此时l和r取得最小值
		综合三种情况都是return -1
	*/
	return -1
}

/*
[labuladong二分法讲解 ](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie)
使用左闭右开区间l<r,并且r=len(num)-1
*/
func Search3(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		}
	}
	/*
		结束条件left == right,l,r的取值范围是[0,len(nums)]
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 15}, expect: -1}求最大值,r=5,l=5
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 4}, expect: -1}中间值,r=3,l=3
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: 12}, expect: 5}最大值,r=5,l=5
		使用{nums: []int{-1, 0, 3, 5, 9, 12}, target: -2}, expect: -1}求最小值,r=0,l=0
		r:[0,len(nums)-1]
		l:[0,len(nums)-1]
		运行到此处时的情况:
		1: l=len(nums)-1,r=len(nums)-1,此时l和r取得最大值,nums[l]!=target
		2: l=r并且l和r都是有效下标,但是nums[l]!=target,nums[r]!=target
		3: 因为是左开右闭区间,所以nums[len(nums)-1]没有被搜索到就结束循环,需要判断nums[len(nums)-1]是否和target相等
		4:l=0,r=0并且nums[0]!=target,此时l和r取得最小值
		综合三种情况得出需要打补丁
	*/
	if nums[l] == target {
		return l
	}
	return -1
}

/*
[labuladong二分法讲解 ](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie)
寻找左侧边界的二分搜索,使用左闭右开区间l<r,并且r=len(num)
*/
func LeftBound1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		}
	}
	/*
		while 终止的条件是 left == right,l的意义代表nums 中小于 target 的元素的个数
		结束条件left == right,l,r的取值范围是[0,len(nums)]
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},求最小值,l=0,r=0
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 0},最小值,r=0,l=0
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 4},中间值,r=4,l=4
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},中间值,r=4,l=4
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 6},中间值,r=6,l=6
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},最大值,r=8,l=8

		所以l和r的取值范围为: r:[0,len(nums)],l:[0,len(nums)]
		运行到此处时的情况:
		1:l=0,r=0并且nums[0]!=target,此时l和r取得最小值
		2:l=0,r=0并且nums[0]==target,此时l和r取得最小值
		3:l=r并且l和r都是有效下标,并且nums[l]==target
		4:l=r并且l和r都是有效下标,并且nums[l]!=target
		5:l=r并且l和r都是有效下标,并且nums[l]==target
		6:l=r=len(nums)
		综合三种情况得出需要打补丁
	*/
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

/*
[labuladong二分法讲解 ](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie)
寻找左侧边界的二分搜索,使用左闭右开区间l<r,并且r=len(num)
*/
func LeftBound2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	/*
		while 终止的条件是 left == right,l的意义代表nums 中小于 target 的元素的个数
		结束条件left == right,l,r的取值范围是[0,len(nums)]
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},求最小值,l=0,r=-1
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 0},最小值,l=0,r=-1
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 4},中间值,l=4,r=3
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},中间值,l=4,r=3
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 6},中间值,l=6,r=5
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},最大值,l=8,r=7

		所以l和r的取值范围为: r:[-1,len(nums)-1],l:[0,len(nums)]
		运行到此处时的情况:
		1:l=0,r=-1并且nums[l]!=target,此时l和r取得最小值
		2:l=0,r=-1并且nums[0]==target,此时l和r取得最小值
		3:l=r+1并且l和r都是有效下标,并且nums[l]==target
		4:l=r+1并且l和r都是有效下标,并且nums[l]!=target
		5:l=r+1并且l和r都是有效下标,并且nums[l]==target
		6:l=r+1,l==len(nums)
		综合三种情况得出需要打补丁
	*/
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

//支持搜索右侧边界的算法
func RightBound1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		}
	}
	/*
		while 终止的条件是 left == right
		结束条件left == right,l,r的取值范围是[0,len(nums)],因为
		if (nums[mid] == target) {
			left = mid + 1;// 这样想: mid = left - 1
		}
		while 循环结束时，nums[left] 一定不等于 target 了，而 nums[left-1] 可能是 target

		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},最小值l=0,r=0
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 1},中间值,l=2,r=2
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 5},中间值,l=6,r=6
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},中间值,l=4,r=4
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 7},最大值,l=8,r=8
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},最大值l=8,r=8
		所以l和r的取值范围为: r:[0,len(nums)],l:[0,len(nums)]
		运行到此处时的情况:
		1:l=0,r=0并且nums[l]!=target,此时l和r取得最小值
		2:l=2,r=2并且nums[l-1]==target
		3:l=r并且l和r都是有效下标,并且nums[l-1]==target
		4:l=r并且l和r都是有效下标,并且nums[l-1]!=target
		5:l=r,l==len(nums)
		6:l=r,l==len(nums)
		综合三种情况得出需要打补丁
	*/
	if l < 1 || nums[l-1] != target {
		return -1
	}
	return l - 1
}

func RightBound2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			l = mid + 1
		}
	}
	/*
		结束条件left == right+1,l,r的取值范围是[0,len(nums)-1],因为
		if (nums[mid] == target) {
			left = mid + 1;// 这样想: mid = left - 1
		}
		while 循环结束时，nums[left] 一定不等于 target 了，而 nums[left-1] 可能是 target

		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},最小值l=0,r=-1
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 1},中间值,l=2,r=1
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 5},中间值,l=6,r=5
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},中间值,l=4,r=3
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 7},最大值,l=8,r=7
		使用{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},最大值l=8,r=7
		所以l和r的取值范围为: r:[-1,len(nums)-1],l:[0,len(nums)]
		运行到此处时的情况:
		1:l=0,r=-1并且nums[r]!=target,此时l和r取得最小值
		2:l=2,r=1并且nums[r]==target
		3:l=r+1并且l和r都是有效下标,并且nums[r]==target
		4:l=r+1并且l和r都是有效下标,并且nums[r]!=target
		5:l=r+1,l==len(nums)并且nums[r]==target
		6:l=r+1,l==len(nums)并且nums[r]!=target
		综合三种情况得出需要打补丁
	*/
	if r < 0 || nums[r] != target {
		return -1
	}
	return r
}
