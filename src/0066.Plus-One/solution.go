package Solution

//overflow可以为任意数字,不限于1
func PlusOne1(digits []int) []int {
	overflow := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + overflow
		if tmp > 9 {
			digits[i] = tmp % 10
			overflow = tmp / 10
		} else {
			overflow = 0
			digits[i] = tmp
		}
	}
	res := digits
	if overflow > 0 {
		res = append([]int{overflow}, res...)
	}
	return res
}

func PlusOne2(digits []int) []int {
	l := len(digits)
	if l < 1 {
		return []int{1}
	}
	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func plusOneAtIndex(nums []int, index int) []int {
	if index < 0 {
		return []int{1}
	}
	if nums[index] < 9 {
		nums[index]++
		return nums
	}
	nums[index] = 0
	return plusOneAtIndex(nums, index-1)
}

// [windliang 解法一 递归](https://leetcode.wang/leetCode-66-Plus-One.html)
func PlusOne3(digits []int) []int {
	return plusOneAtIndex(digits, len(digits)-1)
}

// [windliang 解法二 迭代](https://leetcode.wang/leetCode-66-Plus-One.html)
func PlusOne4(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			break
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
