package Solution

/*
暴力解法
*/
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1)) //用来存储结果
	for i := range nums1 {
		res[i] = -1    //都初始化为-1
		start := false //start为true时表示已经在nums2中找到nums1[i]
		for j := range nums2 {
			if nums2[j] == nums1[i] {
				start = true
			}
			if start && nums2[j] > nums1[i] {
				res[i] = nums2[j]
				break
			}
		}
	}
	return res
}

/*
[leetcode 题解 从右往左单调递减栈](https://leetcode-cn.com/problems/next-greater-element-i/solution/dan-diao-zhan-zong-jie-by-wu-xian-sen-2/)
*/
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	find := make(map[int]int, len(nums2)) //find表示nums2的单调栈结果
	stack := []int{}
	//从右往左求出nums2的单调递减栈
	for i := len(nums2) - 1; i >= 0; i-- {
		/*
			因为是倒着入栈,所以这个for循环实际上是把从元素nums2[i]开始往右,不比nums2[i]大的元素都弹出,
			直到找到比nums2[i]大的元素,同时nums2[i]也要入栈,这些元素放在栈中也没有意义
		*/
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		temp := -1
		if len(stack) != 0 {
			temp = stack[len(stack)-1]
		}
		find[nums2[i]] = temp
		stack = append(stack, nums2[i])
	}
	//计算结果
	res := make([]int, len(nums1))
	for i := range nums1 {
		res[i] = find[nums1[i]]
	}
	return res
}

/*
[leetcode 题解 从左往右单调递减栈](https://leetcode-cn.com/problems/next-greater-element-i/solution/dan-diao-zhan-zong-jie-by-wu-xian-sen-2/)
*/
func nextGreaterElement3(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums2)) //记录nums2中找到的每个元素右边第一个比自己大的值
	var stack []int                    //栈里面最终存储的是右边不存在比自己大的值的元素
	for _, v := range nums2 {
		for len(stack) != 0 && v > stack[len(stack)-1] {
			// 发现有更大的数字，给其下一个更大数字赋值
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for k, v := range nums1 {
		if value, ok := m[v]; ok {
			nums1[k] = value
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}
