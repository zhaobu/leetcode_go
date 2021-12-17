/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start

// 解法1:使用一个临时数组,双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	newNums := make([]int, len(nums1))
	copy(newNums, nums1) //复制一份nums1,方便合并2个数组到nums1中
	var (
		i = 0 //标记newNums下标
		j = 0 //标记nums2下标
		k = 0 //标记nums1下标
	)
	for i < m && j < n {
		if newNums[i] <= nums2[j] {
			nums1[k] = newNums[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	//合并剩下的
	for i < m {
		nums1[k] = newNums[i]
		i++
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

//解法2:从后往前移动
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	var (
		i = m - 1     //nums1要移动的下一个数下标
		j = n - 1     //nums2要移动的下一个数下标
		k = m + n - 1 //nums1下一个数插入的位置下标
	)
	for j >= 0 { //以nums2全部插入到nums1为准
		if i >= 0 && nums1[i] >= nums2[j] { //注意有可能m=0,则i为负数
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// @lc code=end

