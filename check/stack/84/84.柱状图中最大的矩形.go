package main

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start

// 解法1:暴力求解,超时
func largestRectangleArea1(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	maxArea := 0
	for i, height := range heights {
		//寻找左边第一个比heights[i]<的位置
		left := i - 1 //保存当前柱子左边第一个比当前柱子矮的索引
		for left >= 0 && heights[left] >= height {
			left--
		}
		//寻找右边第一个比heights[i]<的位置
		right := i + 1 //保存当前柱子右边第一个比当前柱子矮的索引
		for right < len(heights) && heights[right] >= height {
			right++
		}
		curArea := (right - left - 1) * height //注意宽度计算
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

// 解法2:类似动态规划
func largestRectangleArea2(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	leftLess := make([]int, len(heights))  //记录i左侧且最近的小于其高度的柱子的下标
	rightLess := make([]int, len(heights)) //记录i右侧且最近的小于其高度的柱子的下标
	leftLess[0] = -1
	rightLess[len(heights)-1] = len(heights)

	for i := 1; i < len(heights); i++ { //从左往右依次计算leftLess
		left := i - 1                                  //从i的左边第一个开始比较
		for left >= 0 && heights[left] >= heights[i] { //跳过i的左边所有>=heights[i]的位置
			/*
			   1. 走到这里说明left位置的高度比i位置的要大,应该找比heights[left]
			   更小的位置去和heights[i]比较
			   2. 比heights[left]更小的第一个位置就是leftLess[left]
			   3. 因为是从左往右计算的leftLess,所以小于i的所有leftLess值已知
			*/
			left = leftLess[left]
		}
		leftLess[i] = left
	}

	for i := len(heights) - 2; i >= 0; i-- { //从右往左依次计算rightLess
		right := i + 1 //从i的右边第一个位置开始比较
		for right < len(heights) && heights[right] >= heights[i] {
			right = rightLess[right]
		}
		rightLess[i] = right
	}

	//计算最大面积
	maxArea := 0
	for i, height := range heights {
		curArea := (rightLess[i] - leftLess[i] - 1) * height
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

//解法3:单调递增栈写法1
func largestRectangleArea3(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	leftLess := make([]int, len(heights))    //记录i左侧且最近的小于其高度的柱子的下标
	rightLess := make([]int, len(heights))   //记录i右侧且最近的小于其高度的柱子的下标
	ascStack := make([]int, 0, len(heights)) //单调递增栈,栈中存储的是heights的下标,使用时通过heights来获取高度
	for i, height := range heights {
		// 如果新的元素较小，那就一直把栈内元素弹出来，直到栈顶比新元素小
		for len(ascStack) > 0 && heights[ascStack[len(ascStack)-1]] >= height {
			stackTop := ascStack[len(ascStack)-1] //(栈顶)出栈元素
			rightLess[stackTop] = i               // 新元素是出栈元素向后找第一个比其小的元素
			ascStack = ascStack[:len(ascStack)-1] // 出栈
		}
		//当栈顶元素小于当前元素时,栈顶元素就是当前元素的左边第一个比其小的元素
		if len(ascStack) > 0 {
			leftLess[i] = ascStack[len(ascStack)-1]
		} else {
			leftLess[i] = -1
		}
		ascStack = append(ascStack, i) //如果新的元素比栈顶元素大，就入栈
	}
	//栈里面剩余的元素需要全部弹出,因为他们还没有求rightLess,
	for len(ascStack) > 0 {
		stackTop := ascStack[len(ascStack)-1] //(栈顶)出栈元素
		rightLess[stackTop] = len(heights)    // 新元素是出栈元素向后找第一个比其小的元素
		ascStack = ascStack[:len(ascStack)-1] // 出栈
	}
	//计算最大面积
	maxArea := 0
	for i, height := range heights {
		curArea := (rightLess[i] - leftLess[i] - 1) * height
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

//解法4:单调递增栈写法2
func largestRectangleArea(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	ascStack := make([]int, 0, len(heights)) //单调递增栈,栈中存储的是heights的下标,使用时通过heights来获取高度
	maxArea := 0                             //记录最大面积
	heights = append(heights, 0)             //在heights最后插入一个比谁都小的0,从而在0入栈后会计算完所有的情况
	right := 0                               //计算maxArea时右边界
	left := 0                                //计算maxArea时左边界

	for i, height := range heights {
		// 如果新的元素较小，那就一直把栈内元素弹出来，直到栈顶比新元素小
		for len(ascStack) > 0 && heights[ascStack[len(ascStack)-1]] >= height {
			stackTop := ascStack[len(ascStack)-1] //(栈顶)出栈元素
			right = i                             //当元素出栈时，说明这个新元素是出栈元素向后找第一个比其小的元素
			ascStack = ascStack[:len(ascStack)-1] // 出栈
			if len(ascStack) > 0 {                //当元素出栈后，说明新栈顶元素是出栈元素向前找第一个比其小的元素
				left = ascStack[len(ascStack)-1]
			} else {
				left = -1
			}
			//计算以heights[stackTop]为高度的面积
			curArea := (right - left - 1) * heights[stackTop]
			if curArea > maxArea {
				maxArea = curArea
			}
		}
		ascStack = append(ascStack, i) //如果新的元素比栈顶元素大，就入栈
	}
	return maxArea
}

// @lc code=end
