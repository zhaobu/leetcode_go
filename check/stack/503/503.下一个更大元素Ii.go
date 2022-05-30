package main

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
/*
解法1 暴力法
*/
func nextGreaterElements1(nums []int) []int {
	n := len(nums)

	ret := make([]int, n)
	for i := 0; i < len(nums); i++ {
		ret[i] = -1
		for j := (i + 1) % n; j != i; j = (j + 1) % n {
			if nums[j] > nums[i] {
				ret[i] = nums[j]
				break
			}
		}
	}
	return ret
}

/*
解法2 单调栈
每个元素都进栈一次,在遇到比栈顶元素更大元素时才出栈
*/

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	stack := make([]int, 0, n)

	for i := range ret {
		ret[i] = -1
	}
	/*
		[0,1,2,3]+[0,1,2]
	*/
	for i := 0; i < 2*n-1; i++ {
		j := i % n
		// fmt.Printf("\nnums[%d]=%d,stack=%+v\n", j, nums[j], stack)
		for len(stack) > 0 && nums[j] > nums[stack[len(stack)-1]] {
			// fmt.Printf("nums[%d]=%d > stack=%+v,last=%d\n", j, nums[j], stack, nums[stack[len(stack)-1]])
			top := stack[len(stack)-1]
			ret[top] = nums[j]
			stack = stack[:len(stack)-1]
		}
		/*
			不加这个判断也可以使结果正确,不过其实每个元素只需要入栈一次即可,
			不加判断后入栈多次,也会导致出栈多次
		*/
		if i < n {
			stack = append(stack, j)
		}
	}
	return ret
}

// @lc code=end
