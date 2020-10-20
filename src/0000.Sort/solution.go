package Solution

/*
冒泡排序
*/
func BubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		sort := true
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				sort = false //发生了交换说明还不是有序
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if sort {
			break
		}
	}
	return nums
}

/*
选择排序
*/
func SelectionSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {

	}
	return nums
}
