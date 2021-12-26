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

//  类似动态动态规划,预计算每个位置左右2边最大值
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
func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	res := 0 //记录雨水总量
	st := make([]int, 0, len(height))
	for i := 0; i < len(height); i++ {
		//如果元素>栈顶元素,当前元素是栈顶元素的右边界
		for len(st) > 0 && height[i] > height[st[len(st)-1]] {
			stackTop := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			leftMax := height[st[len(st)-1]]
			width := i - st[len(st)-1] - 1
			if height[i] < leftMax {
				res += (height[i] - height[stackTop]) * width
			} else {
				res += (leftMax - height[stackTop]) * width
			}
		}
		st = append(st, i)
	}
	return res
}

// @lc code=end
