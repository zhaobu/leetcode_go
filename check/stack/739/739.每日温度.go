package main

import "fmt"

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

/*
 解法1
 利用单调递减栈,
 1. 当栈顶元素出栈时,说明栈顶元素遇到了右边第一个更大的元素
 2. 当元素入栈时,说明栈顶元素就是当前元素左边第一个更大的元素
*/
// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	var ret []int
	if len(temperatures) < 1 {
		return ret
	}
	ret = make([]int, len(temperatures))
	descStack := make([]int, 0, len(temperatures))

	for i, v := range temperatures {
		fmt.Printf("i=%d, 入栈前 descStack=%+v\n", i, descStack)
		for len(descStack) > 0 && temperatures[descStack[len(descStack)-1]] < v {
			//比栈顶元素大的下一个元素就是当前遍历到的元素
			top := descStack[len(descStack)-1]
			ret[top] = i - top
			descStack = descStack[:len(descStack)-1]
		}
		fmt.Printf("i=%d, 入栈后 descStack=%+v\n\n", i, descStack)
		descStack = append(descStack, i)
	}
	//遍历完后,栈中剩下的都是单调递减的元素,也就是右边没有比他更高温度的时间
	fmt.Printf("descStack=%+v\n", descStack)
	return ret
}

// @lc code=end
