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
	record := make(map[int]int, n) //用来统计每一个累加和出现的次数
	record[0] = 1
	preSum := make([]int, n+1) //前缀和
	sum := 0                   //记录累加和
	for i, v := range nums {
		sum += v
		preSum[i+1] = sum
		/*
			1. preSum[i+1]表示以nums[i]结尾的前缀和
			2. 假设[j,i]之间的和等于k,也就是 preSum[i+1]-preSum[j+1]=k
			3. 用record记录preSum[j+1]的个数,可以方便求出所有满足第2条的j的数量
		*/
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
