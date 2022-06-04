package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start

/*
解法1 转换为字符数组判断
*/
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	arr := []int{}
	for x > 0 {
		arr = append(arr, x%10)
		x /= 10
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

/*
解法2 反转一半
*/
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	//如果x>0且第一位为0.必定不为回文数
	if x < 0 || x%10 == 0 {
		return false
	}

	/*
		1. 奇数位数字 12321
		2. 偶数位数字 321123
	*/
	half := 0
	for x > half {
		half = half*10 + x%10
		x /= 10
	}

	/*
		1. 奇数位的回文数
		12321: x=12, half =123
		32123: x=32, half =321
		2. 偶数位的回文数
		123321: x=123, half =123
		321123: x=321, half =321

		3. 奇数位的非回文数
		12324: x=12, half =423
		52123: x=52, half =321
		4. 偶数位的非回文数
		123329: x=123, half =923
		921121: x=92, half =1211
	*/
	return half == x || x == half/10
}

// @lc code=end
