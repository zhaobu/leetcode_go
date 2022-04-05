package main

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start

/*
解法1 前缀和暴力法
*/

func subarraySum1(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)

	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	ret := 0
	for i := 0; i < len(preSum); i++ {
		for j := i + 1; j < len(preSum); j++ {
			if preSum[j]-preSum[i] == k {
				ret++
			}
		}
	}

	return ret
}

/*
解法2 前缀和优化
*/
func subarraySum(nums []int, k int) int {
	n := len(nums)

	ret := 0
	record := map[int]int{0: 1} //用来统计每一个累加和出现的次数
	preSum := make([]int, n+1)  //前缀和
	sum := 0                    //记录累加和
	for i, v := range nums {
		sum += v
		preSum[i+1] = sum

		if count, ok := record[sum-k]; ok {
			// fmt.Printf("record=%+v,preSum=%+v,count=%d,sum-k=%d\n", record, preSum, count, sum-k)
			ret += count
		}
		record[sum] += 1 //前缀和为sum的次数累加
	}

	// fmt.Printf("nums=%+v,preSum=%+v\n", nums, preSum)
	// for _, v := range preSum {
	// }
	return ret
}

// @lc code=end
