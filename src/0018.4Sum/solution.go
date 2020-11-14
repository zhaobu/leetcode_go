package Solution

import "sort"

/*
[leetcode 官方方法一：排序 + 双指针 ](https://leetcode-cn.com/problems/4sum/solution/si-shu-zhi-he-by-leetcode-solution/)
*/
func fourSum1(nums []int, target int) [][]int {
	// 数组排序
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0) //存储结果
	//在确定第一个数之后，如果 nums[i]+nums[i+1]+nums[i+2]+nums[i+3]>target，说明此时剩下的三个数无论取什么值，四数之和一定大于 target，因此退出第一重循环；
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		/*
			1:去掉重复的nums[i]
			2:在确定第一个数之后，如果 nums[i]+nums[n−3]+nums[n−2]+nums[n−1]<target，说明此时剩下的三个数无论取什么值，四数之和一定小于target，因此第一重循环直接进入下一轮，枚举 nums[i+1]；
		*/
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		//在确定前两个数之后，如果 nums[i]+nums[j]+nums[j+1]+nums[j+2]>target，说明此时剩下的两个数无论取什么值，四数之和一定大于target，因此退出第二重循环；
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			/*
				1:去掉重复的nums[j]
				2:在确定前两个数之后，如果 nums[i]+nums[j]+nums[n-2]+nums[n-1] < target，说明此时剩下的两个数无论取什么值，四数之和一定小于target，因此第二重循环直接进入下一轮，枚举 nums[j+1]；
			*/
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			//使用双指针遍历第三第四个元素
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					//去除重复的第三个元素
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					//去除重复的第四个元素
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
