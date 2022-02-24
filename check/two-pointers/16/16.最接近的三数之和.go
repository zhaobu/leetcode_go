package main

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	return 0
}

// @lc code=end

/*
题目为 [面试题 16.16. 部分排序](https://leetcode-cn.com/problems/sub-sort-lcci/)
输入： [1,2,4,  7,10,11,7,12,6,7,  16,18,19]
输出： [3,9]

输入： [0]
输出： [-1,-1]

输入： [1,3,5,7,9]
输出：[-1,-1]

*/
func subSort(array []int) []int {
	ret := []int{-1, -1}
	if len(array) < 2 {
		return ret
	}
	/*
	  1.从左向右找到最后一个逆序对的位置 rightIdx
	  2.也就是 rightIdx 的左边存在比 array[array] 大的元素 leftMax
	  3.在rightIdx的右边不存在满足1和2的位置
	*/
	leftMax := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] >= leftMax {
			leftMax = array[i]
		} else {
			ret[1] = i
		}
	}

	/*
	  1.从右向左找到第一个逆序对的位置 leftIdx
	  2.也就是 leftIdx 的右边存在比 array[leftIdx] 小的元素 rightMin
	  3.在 leftIdx 的左边不存在满足1和2的位置
	*/
	rightMin := array[len(array)-1]
	for i := len(array) - 2; i >= 0; i-- {
		if array[i] <= rightMin {
			rightMin = array[i]
		} else {
			ret[0] = i
		}
	}

	return ret
}
