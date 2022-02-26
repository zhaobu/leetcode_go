package main

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
// @lc code=start

/*
 解法3
 可以看成是暴力算法的优化
 1.用三个变量,i,j,maxIdx分别表示滑动窗口的左,右边界和窗口内的最大值下标
 2.移动窗口到下一个位置,如果上一个窗口的最大值还在当前窗口内,用新进到窗口的元素和上一个窗口的最大值比较.
 如果不在,从i遍历到j,求得新窗口的最大值
*/
func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	if len(nums) < 1 {
		return ret
	}
	if k == 1 {
		return nums
	}
	maxIdx := -1   //滑动窗口内最大值的下标
	i := 0         //滑动窗口左边界
	j := i + k - 1 //滑动窗口右边界
	for ; j < len(nums); i, j = i+1, j+1 {
		//如果上一个窗口内的最大值下标还在当前窗口内,就和当前窗口的最后一个元素比较大小
		if maxIdx >= i {
			if nums[j] >= nums[maxIdx] {
				maxIdx = j
			}
		} else {
			//如果不在,就重新遍历求最大值的下标
			maxIdx = i
			for k = i + 1; k <= j; k++ {
				if nums[k] >= nums[maxIdx] {
					maxIdx = k
				}
			}
		}
		ret = append(ret, nums[maxIdx])
	}

	return ret
}

/*
 解法2
 滑动窗口写法2

*/

func maxSlidingWindow2(nums []int, k int) []int {
	ret := []int{}
	if len(nums) < 1 {
		return ret
	}
	if k == 1 {
		return nums
	}

	//单调递减双端队列,存储的是nums数组的下标
	descDeque := make([]int, 0, k)

	// i表示滑动窗口的最后一个元素下标
	for i := 0; i < len(nums); i++ {

		//从队尾弹出所有比nums[i]要小的元素,维护单调递减队列
		for len(descDeque) > 0 && nums[descDeque[len(descDeque)-1]] < nums[i] {
			descDeque = descDeque[:len(descDeque)-1]
		}
		descDeque = append(descDeque, i)
		// fmt.Printf("descDeque=%+v\n", descDeque)
		// 窗口的大小小于k时
		if i < k-1 {
			continue
		}
		//单调递减队列的队头元素就是窗口的最大值
		ret = append(ret, nums[descDeque[0]])
		// 如果窗口向右滑动后,队头元素超出范围就要去掉
		if descDeque[0] <= i-k+1 {
			descDeque = descDeque[1:]
		}
	}

	return ret
}

/*
 解法1
 滑动窗口写法1
*/
// @lc code=start
func maxSlidingWindow1(nums []int, k int) []int {
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
