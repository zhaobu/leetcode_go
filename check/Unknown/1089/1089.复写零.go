package main

/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

// @lc code=start

/*
解法1 2遍遍历
*/
func duplicateZeros(arr []int) {
	if len(arr) == 1 {
		return
	}
	/*
		1. lastZero 表示最后一个可以复写的0在原数组的下标
		2. slow 表示复写0完成后最后一个元素在原数组的下标
	*/
	lastZero := 0
	fast, slow := 0, 0
	for fast < len(arr) {
		if arr[slow] == 0 {
			fast++ //遇到0还需要再占用一个位置
			if fast < len(arr) {
				//如果这个0复写完成后不超过数组长度说明这个0可以复习
				lastZero = slow
			}
		}
		slow, fast = slow+1, fast+1
	}

	if fast == len(arr) {
		slow, fast = slow-1, fast-1
	} else if fast == len(arr)+1 {
		slow, fast = slow-1, fast-2
	}

	// fmt.Printf("slow=%d,fast=%d\n", slow, fast)
	/*
		1. 从后往前开始移动每一个元素
	*/
	for ; slow >= 0; slow, fast = slow-1, fast-1 {
		arr[fast] = arr[slow]
		if arr[slow] == 0 && slow <= lastZero {
			//只有在下标 <= lastZero的0才能复写,否则就是超出范围的0,不用复写
			fast--
			arr[fast] = arr[slow]
		}
	}
	return
}

// @lc code=end
