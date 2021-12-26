package main

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) < 1 || k < 1 {
		return nil
	}

	descDqueue := make([]int, 0, k) //单调递减双端队列

	//队列插入新的位置i处的元素
	push := func(i int) {
		//如果队头不在窗口中就要去除队头元素
		if len(descDqueue) > 0 && descDqueue[0] < i-k+1 {
			descDqueue = descDqueue[1:]
		}

		//在队尾添加元素，要把前面比自己小的元素都删掉，直到遇到更大的元素才停止删除
		for len(descDqueue) > 0 && nums[descDqueue[len(descDqueue)-1]] <= nums[i] {
			descDqueue = descDqueue[:len(descDqueue)-1]
		}
		//元素入队
		descDqueue = append(descDqueue, i)
	}

	//前k个元素先入队
	for i := 0; i < k; i++ {
		push(i)
	}

	res := make([]int, 0, k)
	res = append(res, nums[descDqueue[0]]) //加入窗口在第一个位置时的最大值
	//从第一个位置的后一个位置开水,直到nums的最后一个位置,依次计算结果
	for i := k; i < len(nums); i++ {
		push(i)
		res = append(res, nums[descDqueue[0]])
	}

	return res
}

// @lc code=end
