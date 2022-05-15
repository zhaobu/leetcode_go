package main

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start

/*
解法1 hash表暴力法(超时)
*/
func findLength1(nums1 []int, nums2 []int) int {
	/*
		arr1 记录nums1中每种元素的所有下标
		arr2 记录nums2中每种元素的所有下标
	*/

	arr1, arr2 := map[int][]int{}, map[int][]int{}

	for i, v := range nums1 {
		arr1[v] = append(arr1[v], i)
	}

	for i, v := range nums2 {
		arr2[v] = append(arr2[v], i)
	}
	// fmt.Printf("arr1=%+v,arr2=%+v\n", arr1, arr2)
	n1, n2 := len(nums1), len(nums2)
	ret := 0
	for value, indexs1 := range arr1 {
		if indexs2, ok := arr2[value]; ok {
			for _, index1 := range indexs1 {
				for _, index2 := range indexs2 {
					// fmt.Printf("value=%+v, i=%+v, j=%+v \n", value, i, j)
					count := 0
					for i, j := index1, index2; i < n1 && j < n2 && (nums1[i] == nums2[j]); i, j = i+1, j+1 {
						count++
					}
					if count > ret {
						ret = count
					}
				}
			}
		}
	}

	return ret
}

/*
解法2 动态规划
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
func findLength(nums1 []int, nums2 []int) int {
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
