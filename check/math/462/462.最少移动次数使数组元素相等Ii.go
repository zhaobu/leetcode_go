package main

import (
	"math/rand"
	"sort"
	"time"
)

/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start

/*
解法1 排序
1. 因为排序后对求满足到所有数的距离之和最小的数不影响,所以考虑排序
2. 考虑一下区间[0, len(nums)-1]. 如果只有2个元素,那只要在这2个元素之间的数都满足条件
因为2个点之间的点到左右两边点的距离和就是2个点之间的距离
3. 可以考虑把nums分成多个区间,则满足条件的点就是包含在所有区间之间的点即可
4. 奇数个元素时,最中间区间为一个元素len(nums)/2,偶数个元素时,最中间区间为
[len(nums)/2-1, len(nums)/2]
5. 排序后下标为len(nums)/2的元素满足到所有数的距离之和最小
*/
func minMoves21(nums []int) int {
	sort.Ints(nums)
	pivot := nums[len(nums)/2]
	ret := 0
	for i := range nums {
		if nums[i] > pivot {
			ret += nums[i] - pivot
		} else {
			ret += pivot - nums[i]
		}
	}
	return ret
}

/*
解法2 快速排序
1. 排序后下标为len(nums)/2的元素满足到所有数的距离之和最小
2. 利用快速排序思想找出数组中第len(nums)/2+1小的元素
*/
func minMoves2(nums []int) int {
	target := len(nums) / 2 //目标是找到排序后下标为target的元素

	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	partition := func(left, right int) int {
		//在[i,j]范围内随机选一个下标
		k := rand.Intn(right-left+1) + left //[0,right-left+1)+left
		//选取后的下标元素和最后一个元素交换,分区点被交换到最后一个位置
		nums[k], nums[right] = nums[right], nums[k]
		pivot := nums[right]
		i := left
		for j := left; j < right; j++ {
			if nums[j] <= pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		/*
			1. 把分区元素放到正确位置上
			2. i位置所在的元素就是第一个大于pivot的元素
		*/
		nums[i], nums[right] = nums[right], nums[i]
		return i
	}

	i, j := 0, len(nums)-1
	for i <= j {
		k := partition(i, j) //本次二分后,分区元素所在的下标
		if k == target {
			break
		} else if k < target { //如果分区点元素小于target,说明目标target应该比本次分区点所在元素更大
			i = k + 1
		} else {
			j = k - 1
		}
	}

	pivot := nums[target] //使所有数组元素相等需要的最少移动数所在的位置
	ret := 0
	//求出移动数
	for i := range nums {
		if nums[i] > pivot {
			ret += nums[i] - pivot
		} else {
			ret += pivot - nums[i]
		}
	}
	return ret
}

// @lc code=end
