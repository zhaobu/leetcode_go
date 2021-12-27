package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start

//暴力法
func trap1(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	res := 0
	for i, v := range height {
		//向左求最高点
		leftMax := v
		for j := i - 1; j >= 0; j-- {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}
		//向右求最高点
		rightMax := v
		for k := i + 1; k < len(height); k++ {
			if height[k] > rightMax {
				rightMax = height[k]
			}
		}
		//只有当左右2边都存在更高元素时才计算
		if leftMax > v && rightMax > v {
			//用左右最高点中的较小值作为高累加
			if leftMax < rightMax {
				res += leftMax - v
			} else {
				res += rightMax - v
			}
		}
	}
	return res
}

//动态动态规划,预计算每个位置左右2边最大值
func trap2(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	leftMax := make([]int, len(height))  //计算每个位置左边最大位置,包括每个位置自己
	rightMax := make([]int, len(height)) //计算每个位置右边最大位置,包括每个位置自己
	leftMax[0] = height[0]               //0位置包括自己在内的最大值就是height[0]
	//从左往右计算每个位置左边的最大值
	for i := 1; i < len(height); i++ {
		if leftMax[i-1] > height[i] {
			leftMax[i] = leftMax[i-1]
		} else {
			leftMax[i] = height[i]
		}
	}

	rightMax[len(height)-1] = height[len(height)-1] //len(height)-1位置包括自己在内的最大值就是height[len(height)-1]
	//从右往左计算每个位置的右边的最大值
	for i := len(height) - 2; i >= 0; i-- {
		if rightMax[i+1] > height[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = height[i]
		}
	}
	res := 0
	//根据leftMax和rightMax计算每个位置
	for i := 1; i < len(height)-1; i++ {
		if leftMax[i] < rightMax[i] {
			res += leftMax[i] - height[i]
		} else {
			res += rightMax[i] - height[i]
		}
	}
	return res
}

//用单调递减栈
func trap3(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	res := 0                          //记录雨水总量
	st := make([]int, 0, len(height)) //单调栈存储的是下标，满足从栈底到栈顶的下标对应的数组height中的元素递减
	for i, h := range height {
		//如果元素>栈顶元素,当前元素是栈顶元素的右边界
		for len(st) > 0 && h > height[st[len(st)-1]] {
			stackTop := st[len(st)-1]
			st = st[:len(st)-1] //出栈,然后新栈顶元素就是左边界
			if len(st) == 0 {   //栈为空说明,左边已经没有左边界了
				break
			}
			leftIndex := st[len(st)-1] //新栈顶元素就是左边界
			width := i - leftIndex - 1 //右边界下标i-左边界下标leftIndex-1就是宽度
			if h < height[leftIndex] {
				res += (h - height[stackTop]) * width
			} else {
				res += (height[leftIndex] - height[stackTop]) * width
			}
		}
		st = append(st, i) //每个遍历到的元素都要入栈一次
	}
	return res
}

//双指针解法
func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	/*
		左指针左边的表示已经计算过面积,只会向右移动
		右指针右边的表示已经计算过面积,只会向左移动
	*/
	left, right := 0, len(height)-1
	/*
	 leftMax表示计算到left位置时,左边界到达过的最大高度
	 leftMax表示计算到left位置时,左边界到达过的最大高度
	*/
	leftMax, rightMax := 0, 0
	res := 0 //累加计算结果

	for left < right {
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

// @lc code=end
