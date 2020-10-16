package Solution

/*
[leetcode 官方 方法二：二分查找](https://leetcode-cn.com/problems/sqrtx/solution/x-de-ping-fang-gen-by-leetcode-solution/)
*/
func MySqrt1(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

/*
[二分查找(套用模板,最后返回时打补丁)](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie#san-xun-zhao-you-ce-bian-jie-de-er-fen-cha-zhao)
*/
func MySqrt2(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)>>1
		a := mid * mid
		if a == x {
			return mid
		} else if a < x {
			l = mid + 1
		} else if a > x {
			r = mid - 1
		}
	}
	return r
}

//支持搜索左侧边界的算法
func left_bound(nums []int, target int) int {
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
	//while 终止的条件是 left == right,l的意义代表nums 中小于 target 的元素的个数
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

//支持搜索右侧边界的算法
func right_bound(nums []int, target int) int {
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
	//while 终止的条件是 left == right,l的意义始终表示mid+1,也就是不管有没有找到都是往后一个位置
	if l < 1 || nums[l-1] != target {
		return -1
	}
	return l - 1
}
