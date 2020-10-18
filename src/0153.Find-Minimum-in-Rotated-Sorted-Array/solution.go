package Solution

/*
[windliang 解法一 ](https://leetcode.wang/leetcode-153-Find-Minimum-in-Rotated-Sorted-Array.html)
*/
func FindMin1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func FindMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}

/*
旋转之后有一个特征，就是后面的数字比前面的数字要小，所以只要找出这个后面数字进行返回就可以了；若没有这样的数字即代表没有进行旋转返回数组首元素即可。
[leetcode 官方 方法 1：二分搜索想法](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/solution/xun-zhao-xuan-zhuan-pai-lie-shu-zu-zhong-de-zui-xi/)
*/
func FindMin3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	if nums[r] > nums[0] {
		return nums[0]
	}
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
