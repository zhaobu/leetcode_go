package main

import (
	"fmt"
	"math"
	"runtime/debug"
	"sort"
)

type MySlice []int

func (m MySlice) append1(i int) {
	fmt.Printf(string(debug.Stack()))
	m = append(m, i)
}

func (m *MySlice) append2(i int) {
	fmt.Printf(string(debug.Stack()))
	*m = append(*m, i)
}

func init() {
	fmt.Printf("b=%d\n", 10)
}
func init() {
	fmt.Printf("a=%d\n", 10)
}

func threeSum(nums []int, target int) (ret [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for _, num1 := range nums {
		targetNum := target - num1
		left, right := 0, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == targetNum {
				ret = append(ret, []int{num1, nums[left], nums[right]})
				break
			} else if sum < targetNum {
				left++
			} else {
				right--
			}
		}
	}

	return ret
}

func main() {
	// m := make(MySlice, 2, 4)
	// fmt.Printf("m=%+v\n", m)
	// m.append1(1)
	// m.append2(2)
	// fmt.Printf("m=%+v\n", m)

	// nums := []int{1, 1, 1, 1}
	// target := 3
	// fmt.Printf("%v\n", threeSum(nums, target))

	nums1 := []int{1, 1, 2, 3, 4}
	nums2 := []int{3, 5, 6}
	fmt.Printf("%v\n", merge(nums1, nums2))
}

func merge(nums1, nums2 []int) (ret []int) {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 && n2 == 0 {
		return ret
	}

	i, j := 0, 0
	record := map[int]int{}

	for _, v := range nums1 {
		record[v]++
	}
	for _, v := range nums2 {
		record[v]++
	}

	for i < n1 || j < n2 {
		a, b := math.MaxInt64, math.MaxInt64
		if i < n1 {
			a = nums1[i]
		}
		if j < n2 {
			b = nums2[j]
		}
		c := b
		if a < b {
			c = a
			i++
		} else if a == b {
			j++
			i++
		} else {
			j++
		}
		// fmt.Printf("a=%d, b=%d, c=%d \n", a, b, c)

		if record[c] == 1 {
			ret = append(ret, c)
		}
	}

	return ret
}
