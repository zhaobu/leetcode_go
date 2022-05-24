package main

/*
 * @lc app=leetcode.cn id=556 lang=golang
 *
 * [556] 下一个更大元素 III
 */

// @lc code=start
/*
解法1
做这道题前可以先做 31.[下一个排列](https://leetcode.cn/problems/next-permutation/description/)
*/
func nextGreaterElement(n int) int {

	/*
		1. 如果每个数都比它后一个数大,那这个数一定是最大数
		2. 把数n变为数组
	*/
	check := func(n int) (can bool, nums []int) {
		last := -1
		i := 1
		for n > 0 {
			cur := n % 10
			if !can && cur < last {
				can = true
			}
			i++
			nums = append(nums, cur)
			last = cur
			n /= 10
		}
		return can, nums
	}

	can, nums := check(n)
	if !can {
		return -1
	}

	//反转数组nums
	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse(nums)
	// fmt.Printf("nums=%+v\n", nums)

	//从后往前查找第一个比它后面元素小的元素
	index := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}

	/*
		1. [index+1,len(nums))范围内,元素都是递减的
		2. 找到(index, len(nums))范围内最小的比nums[index]大的元素,假设下标为j, 然后和nums[index]交换
		3. 交换后nums[j+1]<=nums[j]<= nums[j-1]依旧满足[index+1,len(nums))范围内,元素都是递减的
	*/
	for i := len(nums) - 1; i > index; i-- {
		if nums[i] > nums[index] {
			nums[i], nums[index] = nums[index], nums[i]
			break
		}
	}

	//反转index之后的数组
	reverse(nums[index+1:])

	// fmt.Printf("nums=%+v\n", nums)
	num := 0 //记录符合条件的最小整数
	maxNum := 1<<31 - 1
	// fmt.Printf("nums=%+v,maxNum=%d\n", nums, maxNum)
	for i := range nums {
		num = num*10 + nums[i]
		// fmt.Printf("num=%+v,maxNum=%d\n", num, maxNum)
		if num > maxNum {
			return -1
		}
	}

	return num
}

// @lc code=end
