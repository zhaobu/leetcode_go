package main

/*
 * @lc app=leetcode.cn id=2269 lang=golang
 *
 * [2269] 找到一个数字的 K 美丽值
 */

// @lc code=start

/*
解法1
*/
func divisorSubstrings(num int, k int) int {
	m := 0         //num的长度
	arr := []int{} //用切片保存num每个数字
	for n := num; n > 0; n /= 10 {
		m++
		arr = append([]int{n % 10}, arr...)
	}

	//将切片变为数字
	getNum := func(arr []int) (n int) {
		for i := range arr {
			n = n*10 + arr[i]
		}
		return n
	}

	ret := 0
	last := m - k //最后一个可能是美丽数的下标
	for i := 0; i <= last; i++ {
		curNum := getNum(arr[i : i+k])
		if curNum != 0 && num%curNum == 0 {
			ret++
		}
	}

	return ret
}

// @lc code=end
