package Solution

/*[leetcode官方 ](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/solution/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-de-zui--16/)
 */
func FindMin1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}
