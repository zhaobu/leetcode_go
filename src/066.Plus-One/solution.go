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
