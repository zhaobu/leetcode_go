package Solution

import (
	"sort"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

// 找出每个house最佳使用的heater,并且求得house到heater的距离
func find(nums []int, target int) int {
	if len(nums) == 1 {
		return abs(target, nums[0])
	}
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			return 0
		}
	}

	a, b := nums[0], nums[len(nums)-1]
	if l < len(nums) {
		b = nums[l]
	}
	if l > 0 {
		a = nums[l-1]
	}
	return min(abs(a, target), abs(b, target))
}

func FindRadius1(houses []int, heaters []int) int {
	//升序排列heaters
	sort.Slice(heaters, func(i, j int) bool {
		return heaters[i] <= heaters[j]
	})

	res, a := 0, 0
	for _, v := range houses {
		a = find(heaters, v)
		if a > res {
			res = a
		}
	}
	return res
}
