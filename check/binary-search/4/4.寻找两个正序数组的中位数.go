package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start

/*
解法1 二分法
*/
// func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
// 	m, n := len(nums1), len(nums2)

// 	left, right := nums1[0], nums1[m-1]
// 	if nums2[0] < left {
// 		left = nums2[0]
// 	}
// 	if nums2[n-1] > right {
// 		right = nums2[n-1]
// 	}

// 	pairSum := left + sort.Search(right-left, func(sum int) bool {
// 		sum += left
// 		cnt := 0
// 		i, j := 0, n-1
// 		for i < m && j >= 0 {
// 			if nums1[i]+nums2[j] > sum {
// 				j--
// 			} else {
// 				/*
// 					1. 如果nums1[i]+nums2[j]<=sum,则j的取值范围[0,j]之内都满足.
// 					2. 所以对于每一个i,都有j+1个数对满足条件
// 				*/
// 				cnt += j + 1
// 				i++
// 			}
// 		}
// 		return cnt >= k
// 	})
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
