package Solution

import "sort"

//[小浩算法](https://www.geekxh.com/1.0.数组系列/008.html#_02、题目分析)
func ThreeSum1(nums []int) [][]int {
	//以升序排序
	sort.Ints(nums)
	//存储满足条件的结果
	res := [][]int{}
	// 枚举第一个元素
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1         //第二个元素
		k := len(nums) - 1 //第三个元素
		if nums[i] > 0 {   //如果固定下来的数（上面蓝色框框）本身就大于 0，那三数之和必然无法等于 0
			break
		}
		//保证第一个元素不重复
		if i == 0 || nums[i-1] != nums[i] {
			for j < k {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					for j < k {
						j++
						k--
						//只要第二个和第三个元素其中有一个变化了,就要重新判断
						if nums[j] != nums[j-1] || nums[k] != nums[k+1] {
							break
						}
					}
				} else if sum > 0 { //如果和大于0，那就说明 right 的值太大，需要左移
					k--
				} else { //如果和小于0，那就说明 left 的值太小，需要右移
					j++
				}
			}
		}
	}
	return res
}

//[leetcode 官方:方法一](https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/)
func ThreeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
