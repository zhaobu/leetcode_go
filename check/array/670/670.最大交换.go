package main

/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start

/*
解法1 贪心
*/
func maximumSwap1(num int) int {
	arrNum := []int{}
	for num > 0 {
		arrNum = append([]int{num % 10}, arrNum...)
		num /= 10
	}

	//记录每个数字最后出现的下标
	last := [10]int{}
	for i, v := range arrNum {
		last[v] = i
	}
	// fmt.Printf("last=%+v\n", last)
	for i, v := range arrNum {
		for j := 9; j > 0; j-- {
			/*
				1. j > v 表示从后往前找,比当前数字大的最大数字,并且因为从后往前找,当这个数重复时
				能保证是最后一个元素
				2. last[j] > i ,这个最大元素的下标要大于i
			*/
			if last[j] > i && j > v {
				// fmt.Printf("i=%+v,j=%v,v=%v\n", i, j, v)
				arrNum[i], arrNum[last[j]] = arrNum[last[j]], arrNum[i]
				goto flag
			}
		}
	}
flag:
	ret := 0
	for _, v := range arrNum {
		ret = ret*10 + v
	}
	return ret
}

// @lc code=end
