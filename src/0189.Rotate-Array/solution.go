package Solution

// [小浩算法](https://www.geekxh.com/1.0.数组系列/004.html#_03、题目解答)
func Rotate1(nums []int, k int) {
	reverse(nums)
	reverse(nums[0 : k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)-1-i; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

//[leetcode官方 方法1:暴力解法](https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode/)
func Rotate2(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	previous, tmp := 0, 0
	for j := 0; j < k; j++ {
		previous = nums[len(nums)-1]
		for i := 0; i < len(nums); i++ {
			tmp = nums[i]
			nums[i] = previous
			previous = tmp
		}
	}
	return
}

// [leetcode官方 方法2:使用额外的数组](https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode/)
func Rotate3(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		res = append(res, nums[(len(nums)-k%len(nums)+i)%len(nums)])
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
	return
}

//[leetcode官方 方法2:使用环状替换]
func Rotate4(nums []int, k int) {
	n := len(nums)
	k = k % n
	if n < 2 || k == 0 {
		return
	}

	start, tmp := 0, 0
	pos := start
	pre := nums[pos]
	for count := 0; count < n; {
		for {
			pos = (pos + k) % n
			tmp = nums[pos]
			nums[pos] = pre
			pre = tmp
			count++
			if count == n {
				return
			}
			if pos == start {
				break
			}
		}
		start++
		pos = start
		pre = nums[pos]
	}
	return
}
