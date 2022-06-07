package main

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start

/*
解法1 动态规划
*/
func jump1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, n)
	dp[n-1] = 0 //最后一格不用跳
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= n-1-i {
			dp[i] = 1
		} else {
			dp[i] = n         //初始化为n,表示不能到达
			if nums[i] == 0 { //表示不能到达
				continue
			}
			next := i + nums[i]
			for j := i + 1; j <= next && j < n; j++ {
				dp[i] = min(dp[i], 1+dp[j])
			}
		}
	}

	return dp[0]
}

/*
解法2 贪心
*/
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	farthest := 0 //当前能到达的最远距离
	end := 0      //上次跳跃可达范围的最远距离(下次的最右起跳点)
	jumpCnt := 0  //跳跃次数

	for i := 0; i < n-1; i++ {
		//从下标i能到达的最远距离和已记录最远距离比较,取最大值
		cur := nums[i] + i
		if cur > farthest {
			farthest = cur
		}

		//到达上次跳跃能到达的右边界了
		if end == i {
			end = farthest // 目前能跳到的最远位置变成了下次起跳位置的有边界
			jumpCnt++      // 进入下一次跳跃
		}
	}
	return jumpCnt
}

// @lc code=end
