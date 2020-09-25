package Solution

func TwoSum1(nums []int, target int) []int {

	n := len(nums)

	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func TwoSum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok && sub != v {
			return []int{i, j}
		}
	}

	return nil
}

func TwoSum3(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		}
		m[v] = i
	}

	return nil
}
