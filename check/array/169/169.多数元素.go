package main

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start

/*
解法1 map
*/
func majorityElement1(nums []int) int {
	record := map[int]int{}
	ret := 0
	n := len(nums) / 2
	for _, v := range nums {
		record[v]++
		if record[v] > n {
			ret = v
			break
		}
	}

	return ret
}

/*
解法2 摩尔投票法(Boyer-Moore)
*/
func majorityElement2(nums []int) int {
	count := 0
	target := 0
	for _, v := range nums {
		if count == 0 {
			target = v
		}
		if target == v {
			count++
		} else {
			count--
		}
	}
	return target
}

/*
解法3 分治
*/
func majorityElement(nums []int) int {

	//求出num在[i,j]区间上出现的次数
	getCount := func(num, i, j int) (count int) {
		for ; i <= j; i++ {
			if nums[i] == num {
				count++
			}
		}
		return count
	}
	/*
		求[i,j]区间的众数
	*/
	var divid func(i, j int) int
	divid = func(i, j int) int {
		//如果只剩一个元素,众数就是这个数
		if i >= j {
			return nums[i]
		}

		mid := i + (j-i)>>1
		left := divid(i, mid)    //左半部分众数,注意区间是[i,mid],不是[i,mid)
		right := divid(mid+1, j) //右半部分众数

		if left == right {
			return left
		}

		//求出左边众数在整个区间上的次数
		leftCount := getCount(left, i, j)
		rightCount := getCount(right, i, j)

		if leftCount > rightCount {
			return left
		}
		return right
	}

	return divid(0, len(nums)-1)
}

// @lc code=end
