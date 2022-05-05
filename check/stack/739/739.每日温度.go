package main

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start

/*
解法1: 倒推法,有些类似动态规划
i 用来扫描所有的元素，从右往左扫描（ 逐渐递减），一开始i指向倒数第2个元素
对于每一个 ，一开始令 j = i +1
1. 如果 temperatures[i] < temperatures[j] 那么 ret[i] = j-i 然后 i--
2. 如果 ret[j] = 0,说明 j 的右边没有更大的元素,自然也不会有比 temperatures[i] 更大的元素 那么ret[i] = 0 , 然后 i--
3. 否则，设置 j = j + ret[j],回到步骤1
*/
func dailyTemperatures(temperatures []int) []int {
	m := len(temperatures)
	ret := make([]int, m)
	if m == 1 {
		return ret
	}

	for i := m - 2; i >= 0; i-- {
		j := i + 1
		for {
			if temperatures[j] > temperatures[i] {
				ret[i] = j - i
				break
			} else if ret[j] == 0 {
				/*
					1. temperatures[j] <= temperatures[i]
					2. ret[j] == 0 说明右边所有值都 < temperatures[j]
				*/
				ret[i] = 0
				break
			} else {
				/*
					此处没有break,不会死循环,因为 j += ret[j]
					会让j不断变大,总会走到前面2个分支处
				*/
				j += ret[j]
			}
		}
	}
	return ret
}

/*
 解法2
 利用单调递减栈,
 1. 当栈顶元素出栈时,说明栈顶元素遇到了右边第一个更大的元素
 2. 当元素入栈时,说明栈顶元素就是当前元素左边第一个更大的元素
*/
func dailyTemperatures1(temperatures []int) []int {
	m := len(temperatures)
	ret := make([]int, m)
	if m == 1 {
		return ret
	}
	descStack := []int{}

	for i, v := range temperatures {
		// fmt.Printf("i=%d, 入栈前 descStack=%+v\n", i, descStack)
		for len(descStack) > 0 && temperatures[descStack[len(descStack)-1]] < v {
			//比栈顶元素大的下一个元素就是当前遍历到的元素
			top := descStack[len(descStack)-1]
			ret[top] = i - top
			descStack = descStack[:len(descStack)-1]
		}
		// fmt.Printf("i=%d, 入栈后 descStack=%+v\n\n", i, descStack)
		descStack = append(descStack, i)
	}
	//遍历完后,栈中剩下的都是单调递减的元素,也就是右边没有比他更高温度的时间
	// fmt.Printf("descStack=%+v\n", descStack)
	return ret
}

// @lc code=end
