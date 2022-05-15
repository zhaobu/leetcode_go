package main

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start

/*
解法1 滑动窗口
1. 用nums1的第一个元素对齐nums2的每一个下标,并且求得每次对齐时的最长公共子数组,并更新最大值
2. 用nums2的第一个元素对齐nums1的每一个下标,并且求得每次对齐时的最长公共子数组,并更新最大值
3. 在每次对齐时,如果发现已经求得的最长公共子数组长度比剩下的元素还大,那后面的匹配情况就没必要进行了
比如 nums1 = [1,2,3,2,1], num2 = [3,2,1,4,7], 当用nums2的第一个元素匹配到nums1的下标2时已经
求得的最长公共子数组为3,那nums1后面的下标就没必要进行匹配了
*/
func findLength(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)

	ret := 0
	/*
		1. i, j分别表示nums1和nums2的开始下标
		2. nums1和nums2的下标分别从i,j开始,统计最长的公共子数组
	*/
	getMax := func(i, j int) {
		count := 0
		for ; i < n1 && j < n2; i, j = i+1, j+1 {
			if nums1[i] == nums2[j] {
				count++
			} else { //只要nums1和nums2所在元素不等,公共子数组长度就变为0
				count = 0
			}
			if count > ret {
				ret = count
			}
		}
	}

	//分别用nums2的第一个元素对齐nums1的每一个下标
	for i := range nums1 {
		getMax(i, 0)
		/*
			如果nums1剩下的元素个数比已经求得的最长公共子数组长度还短,那后面的匹配就没必要进行了
		*/
		if ret >= n1-i {
			break
		}
	}

	//分别用nums1的第一个元素对齐nums2的每一个下标
	for i := range nums2 {
		getMax(0, i)
		if ret >= n2-i {
			break
		}
	}
	return ret
}

/*
解法2 动态规划
dp[i][j]表示nums1[:i]和nums2[:j]的最长公共后缀
要求的最长公共子数组就转变为在求最长公共后缀的dp过程中更新最大值即可
*/
func findLength2(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)

	/*
		dp[i][j]表示nums1[:i]和nums2[:j]的最长公共后缀
	*/
	dp := make([][]int, n1)

	// fmt.Printf("dp[%d]=%+v\n", 0, dp[0])
	ret := 0
	for i := 0; i < n1; i++ {
		dp[i] = make([]int, n2)
		for j := 0; j < n2; j++ {
			if nums1[i] == nums2[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
		// fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return ret
}

/*
解法3 动态规划优化
*/
func findLength3(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)

	/*
		dp[i][j]表示nums1[:i]和nums2[:j]的最长公共后缀
	*/
	dp := make([]int, n1)
	// fmt.Printf("dp[%d]=%+v\n", 0, dp[0])
	ret := 0
	pre := 0
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			old := dp[j] //提前保存dp[i-1][j-1]
			dp[j] = 0    //必须要初始化一次dp[j],因为二维数组时每次dp[i][j]其实都相当于等于0
			if nums1[i] == nums2[j] {
				if i > 0 && j > 0 {
					dp[j] = pre + 1
				} else {
					dp[j] = 1
				}
			}
			pre = old
			if dp[j] > ret {
				ret = dp[j]
			}
		}
		// fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return ret
}

// @lc code=end
