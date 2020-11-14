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

// [labuladong 一个函数秒杀 2Sum 3Sum 4Sum 问题](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485789&idx=1&sn=efc1167b85011c019e05d2c3db1039e6&chksm=9bd7f755aca07e43405baeac62c76b44d8438fe8a69ae77e87cbb5121e71b6ee46f4c626eb98&scene=21#wechat_redirect)
func ThreeSum3(nums []int) [][]int {
	// 数组排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0) //存储结果

	// 穷举 threeSum 的第一个数
	for first := 0; first < len(nums); first++ {
		// 对 0 - nums[first] 计算 twoSum
		target := 0 - nums[first]
		second, third := first+1, len(nums)-1
		for second < third {
			left, right := nums[second], nums[third]
			sum := left + right
			if sum < target {
				for second < third && nums[second] == left {
					second++
				}
			} else if sum > target {
				for second < third && nums[third] == right {
					third--
				}
			} else {
				res = append(res, []int{nums[first], left, right})
				for second < third && nums[second] == left {
					second++
				}
				for second < third && nums[third] == right {
					third--
				}
			}
		}
		// 跳过第一个数字重复的情况，否则会出现重复结果
		for first < len(nums)-1 && nums[first] == nums[first+1] {
			first++
		}
	}

	return res
}

//[leetcode 官方:方法一 优化版](https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/)
func ThreeSum4(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	//在确定第一个数之后，如果 nums[first]+nums[first+1]+nums[first+2]>0,说明此时剩下的2个数无论取什么值，3数之和一定大于 0，因此退出第一重循环；
	for first := 0; first < n-2 && nums[first]+nums[first+1]+nums[first+2] <= 0; first++ {
		/*
			1:去掉重复的nums[first]
			2:在确定第一个数之后，如果 nums[first]+nums[n-2]+nums[n-1] < 0，说明此时剩下的2个数无论取什么值，3数之和一定小于0，因此第一重循环直接进入下一轮，枚举 nums[first+1]；
		*/
		if first > 0 && nums[first] == nums[first-1] || nums[first]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		//使用双指针遍历第2第3个元素
		for left, right := first+1, n-1; left < right; {
			if sum := nums[first] + nums[left] + nums[right]; sum == 0 {
				res = append(res, []int{nums[first], nums[left], nums[right]})
				//去除重复的第2个元素
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				//去除重复的第3个元素
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
