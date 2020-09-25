package Solution

import "sort"

func Intersect1(nums1 []int, nums2 []int) []int {
	nums0 := map[int]int{}
	for _, v := range nums1 {
		nums0[v] += 1
	}

	k := 0
	for _, v := range nums2 {
		if nums0[v] > 0 {
			nums0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}

func Intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	res := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
