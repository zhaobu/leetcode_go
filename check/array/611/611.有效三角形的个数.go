package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start

/*
解法1 二分法
*/
func triangleNumber1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	/*
	   a,b,c

	*/
	search := func(i, j, sum int) int {
		for i <= j {
			mid := i + (j-i)>>1
			if sum <= nums[mid] {
				j = mid - 1
			} else {
				if mid == n-1 || nums[mid+1] >= sum {
					return mid
				}
				i = mid + 1
			}
		}
		return -1
	}

	ret := 0

	for a := 0; a < n-2; a++ {
		for b := a + 1; b < n-1; b++ {
			i, j := b+1, n-1
			k := search(i, j, nums[a]+nums[b])
			if k != -1 {
				ret += k - b
			}
		}
	}

	return ret
}

/*
解法2 双指针
*/
func triangleNumber(nums []int) (ret int) {
	sort.Ints(nums)
	n := len(nums)

	/*
	   a,b,c分别为三条边
	   a+c > b
	   b+c > a
	   只需要满足a+b > c即可
	*/
	for a := 0; a < n-2; a++ {
		c := a + 1
		for b := a + 1; b < n-1; b++ {
			sum := nums[a] + nums[b]
			for c+1 < n && sum > nums[c+1] {
				c++
			}
			// fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
			if c > b && c < n {
				ret += c - b
			}
		}
	}

	return ret
}

// @lc code=end
