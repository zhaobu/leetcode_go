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

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	getKthElement := func(nums1, nums2 []int, k int) int {
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

	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}

}

/*
解法2 暴力法
不满足时间复杂度为log(m+n)的要求
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	/*
	   0,1,2,3    (4+1)/2=2
	   0,1,2,3,4  (5+1)/2=3

	   0,1,2,3,4,5   (6+1)/2=3
	   0,1,2,3,4,5,6 (7+1)/2=4
	*/

	i, j := 0, 0 //分别表示nums1和nums2的下标
	/*
	   遍历结束时,
	   如果是偶数: left:中位数第一位, right:中位数第二位
	   如果是奇数: left:就是中位数,right:中位数后一位
	*/
	left, right := 0, 0
	count := (m + n + 1) >> 1 //遍历的次数
	for ; count >= 0; count-- {
		left = right
		if i < m && j < n {
			if nums1[i] <= nums2[j] {
				right = nums1[i]
				i++
			} else {
				right = nums2[j]
				j++
			}
		} else {
			if i == m && j < n {
				right = nums2[j]
				j++
			} else if j == n && i < m {
				right = nums1[i]
				i++
			}
		}
		// fmt.Printf("left=%d,right=%d\n", left, right)
	}
	if (m+n)&1 == 1 {
		return float64(left)
	} else {
		return float64(left+right) / 2
	}
}

// @lc code=end
