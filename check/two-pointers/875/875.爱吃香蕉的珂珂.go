package main

/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start

/*
解法1 二分查找
*/
func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	//找到最大值
	maxPie := piles[0]
	for i := 1; i < n; i++ {
		if piles[i] > maxPie {
			maxPie = piles[i]
		}
	}
	if h == n {
		return maxPie
	}

	calTime := func(k int) (ret int) {
		for _, v := range piles {
			ret += v / k
			if v%k != 0 {
				ret += 1
			}
		}
		return ret
	}

	left, right := 1, maxPie
	for left <= right {
		mid := left + (right-left)>>1
		//计算当前速度吃完需要多长时间
		curTime := calTime(mid)
		if curTime > h {
			left = mid + 1
		} else {
			if mid == 1 || calTime(mid-1) > h {
				return mid
			}
			right = mid - 1
		}
	}

	return -1
}

// @lc code=end
