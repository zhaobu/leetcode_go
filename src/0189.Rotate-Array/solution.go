package Solution

/*
[小浩算法](https://www.geekxh.com/1.0.数组系列/004.html#_03、题目解答)
*/
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

func Rotate2(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp = nums[0]
		for j := 0; j < len(nums); j++ {
			nums[j] = nums[j-1]
		}
	}
	return
}
