package main

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start

/*
解法1 二分查找
1. 数组是经过不重复元素升序数组旋转而来,旋转后左半部分元素>右半部分元素,
nums[0]是左半部分最小元素,nums[n-1]是右半部分最大元素
2. pre表示mid的前一个元素, next表示mid的后一个元素
最小元素nums[i]是唯一满足pre>nums[i]<next的元素
3. 如果整个数组旋转n次,就会和原升序数组一样,则
nums[0] < num[n-1], 否则 nums[0] > num[n-1]
4. 如果不满足条件2,一定不是最小元素
如果元素比右半部分最大元素还大肯定是落在左半部分,最小元素在当前元素右边
如果元素比左半部分最小元素还小肯定是落在右半部分,最小元素在当前元素左边
*/
func findMin1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	i, j := 0, n-1

	//恰好旋转n次,回到原点
	if nums[0] < nums[n-1] {
		return nums[0]
	}

	for i <= j {
		mid := i + (j-i)>>1
		pre, next := nums[(mid+n-1)%n], nums[(mid+1)%n]
		/*
			1. pre表示mid的前一个元素, next表示mid的后一个元素
			2. 最小元素nums[i]是唯一满足pre>nums[i]<next的元素
		*/
		if nums[mid] < pre && nums[mid] < next {
			return nums[mid]
		} else if nums[mid] < nums[0] {
			/*
				1. 左半部分元素>右半部分元素
				2. nums[0]是左半部分最小元素,nums[n-1]是右半部分最大元素
				3. 如果元素比左半部分最小元素还小,则mid肯定是落在右半部分,应该往左找
			*/
			j = mid - 1
		} else { //nums[mid] > nums[n-1]
			// 4. 如果元素比右半部分最大元素还大,则mid肯定是落在左半部分,应该往右找
			i = mid + 1
		}

	}

	// 不会走到这里
	return -1
}

/*
解法2 二分查找
和方法1一样, 只不过不用提前判断当前数组是否恰好旋转n次
*/
func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	i, j := 0, n-1

	for i <= j {
		mid := i + (j-i)>>1
		pre, next := nums[(mid+n-1)%n], nums[(mid+1)%n]
		/*
			1. pre表示mid的前一个元素, next表示mid的后一个元素
			2. 最小元素nums[i]是唯一满足pre>nums[i]<next的元素
		*/
		if nums[mid] < pre && nums[mid] < next {
			return nums[mid]
		} else if nums[mid] > nums[n-1] {
			/*
				1. 左半部分元素>右半部分元素
				2. nums[0]是左半部分最小元素,nums[n-1]是右半部分最大元素
				3. 如果元素比右半部分最大元素还大,则mid肯定是落在左半部分,应该往右找
			*/
			i = mid + 1
		} else { //nums[mid] < nums[0]
			// 4. 如果元素比左半部分最小元素还小,则mid肯定是落在右半部分,应该往左找
			j = mid - 1
		}

	}

	// 不会走到这里
	return -1
}

// @lc code=end
