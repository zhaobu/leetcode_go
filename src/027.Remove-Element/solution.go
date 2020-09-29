package Solution

/*
[解法1 双指针](https://leetcode-cn.com/problems/remove-element/solution/yi-chu-yuan-su-by-leetcode/)
*/
func RemoveElement1(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/*
[解法2 双指针——当要删除的元素很少时](https://leetcode-cn.com/problems/remove-element/solution/yi-chu-yuan-su-by-leetcode/)
*/
func RemoveElement2(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}
