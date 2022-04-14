package main

import "fmt"

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start

/*
解法1 利用index标记为负数来当成hash表
*/
func firstMissingPositive1(nums []int) int {
	m := len(nums)
	if m < 1 {
		return 1
	}
	/*
		1. 长度为m的数组最多能存储从[1,m]这些正整数
		如果全部存储的是这些正数,那最小未出现的就是m+1,
		否则最小未出现的正整数就是第一个未存储的正整数
		2. 把不在[1,m]范围内的数变为0
	*/
	fmt.Printf("nums=%+v;", nums)

	for i := 0; i < m; i++ {
		if nums[i] < 1 || nums[i] > m {
			nums[i] = 0
		}
	}

	//求绝对值
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	/*
		1. 如果出现一个正整数x, x范围必定为[1,m]
		2. 用下标x-1来标记x出现过,也就是把nums[x-1]标记为负数
	*/
	for i := 0; i < m; i++ {
		if nums[i] == 0 {
			continue
		}
		/*
		   1. nums[i]有可能被标记,取绝对值后就能得到真实的原始值
		   2. 用abs(nums[i]) - 1求得的就是用来标记nums[i]的下标
		*/
		index := abs(nums[i]) - 1
		if index == m { //
			continue
		}
		if nums[index] > 0 { //如果nums[index]为负数,说明index+1之前已经出现过,不用重复标记
			nums[index] = -nums[index] //标记当前下标对应的元素
		} else if nums[index] == 0 {
			nums[index] = -(m + 1) //如果赋值为在[-m,-1]之间的一个数,有可能会导致误认为是出现过m+1
		}
	}
	fmt.Printf("nums=%+v\n", nums)
	//找到第一个>=0的下标
	i := 0
	for i < m && nums[i] < 0 {
		i++
	}
	return i + 1
}

/*
解法2 利用index标记为负数来当成hash表
优化初始值,让代码更简洁
*/
func firstMissingPositive(nums []int) int {
	m := len(nums)
	if m < 1 {
		return 1
	}
	/*
		1. 长度为m的数组最多能存储从[1,m]这些正整数
		如果全部存储的是这些正数,那最小未出现的就是m+1,
		否则最小未出现的正整数就是第一个未存储的正整数
		2. 把不在[1,m]范围内的数变为m+1
		3. 变为m+1而不变为0或者[1,m]之间的数字,是为了不影响后面求绝对值后影响计算
	*/
	// fmt.Printf("nums=%+v;", nums)

	for i := 0; i < m; i++ {
		if nums[i] < 1 || nums[i] > m {
			nums[i] = m + 1
		}
	}

	//求绝对值
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	/*
		1. 如果出现一个正整数x, x范围必定为[1,m]
		2. 用下标x-1来标记x出现过,也就是把nums[x-1]标记为负数
	*/
	for i := 0; i < m; i++ {
		/*
			   1. nums[i]有可能被标记,取绝对值后就能得到真实的原始值
			   2. 用abs(nums[i]) - 1求得的就是用来标记nums[i]的下标
			   3. 因为不在[1.m]区间的值都初始化为m+1,所以nums[i]范围在[1,m+1]之间
			    abs(nums[i]) - 1范围在[0,m]之间,所以判断时只需要index<m即可过滤出原始值在
				[1,m]之间的那些元素
		*/
		index := abs(nums[i]) - 1
		if index < m && nums[index] > 0 { //如果nums[index]为负数,说明index+1之前已经出现过,不用重复标记
			nums[index] = -nums[index] //标记当前下标对应的元素
		}
	}
	// fmt.Printf("nums=%+v\n", nums)
	//找到第一个>=0的下标
	i := 0
	for i < m && nums[i] < 0 {
		i++
	}
	return i + 1
}

/*
解法3 交换
1. 长度为m的数组最多能存储从[1,m]这些正整数如果全部存储的是这些正数,那最小未出现的就是m+1,
否则最小未出现的正整数就是第一个未存储的正整数
2. 如果nums[i]在[1,m]之间,那就吧他放到下标为nums[i]-1的位置
3. 这样最终,下标i位置上放的数就是i+1. 如果不是这样说明第一个没出现过的正整数就是i+1
*/
func firstMissingPositive3(nums []int) int {
	m := len(nums)
	if m < 1 {
		return 1
	}
	// fmt.Printf("nums=%+v;", nums)

	for i := 0; i < m; {
		if nums[i] < 1 || nums[i] > m { //如果nums[i]不在[1,m]范围内,也不用处理
			i++
			continue
		}
		index := nums[i] - 1        //nums[i]应该放在下标为index的位置上,交换nums[i]和nums[index]
		if nums[index] == nums[i] { // 如果nums[i]已经放在正确的位置上面,就应该去处理下一个i
			i++
			continue
		}
		nums[index], nums[i] = nums[i], nums[index]
	}
	// fmt.Printf("nums=%+v\n", nums)
	i := 0
	for ; i < m && nums[i] == i+1; i++ {
	}
	return i + 1
}

// @lc code=end
