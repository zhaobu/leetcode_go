# [1. Add Sum][title]

## Description

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

## 题意

> 给定一个整数数组 nums  和一个目标值 target，请你在该数组中找出和为目标值的那   两个   整数，并> 返回他们的数组下标。
> 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

## 题解

### 思路 1

> 简单粗暴些，两重循环，遍历所有情况看相加是否等于目标和，如果符合直接输出。

时间复杂度：两层 for 循环，O（n²）
空间复杂度：O（1）

```go
func TwoSum1(nums []int, target int) []int {

	n := len(nums)

	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
```

### 思路 2

> 我们可以把数组的每个元素保存为 hash 的 key，下标保存为 hash 的 value 。
>
> 这样只需判断 sub 在不在 hash 的 key 里就可以了，而此时的时间复杂度仅为 O（1）！
>
> 需要注意的地方是，还需判断找到的元素不是当前元素，因为题目里讲一个元素只能用一次。

```go
func TwoSum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}
```

### 思路 3

> 看解法二中，两个 for 循环，他们长的一样，我们当然可以把它合起来。复杂度上不会带来什么变化，变化仅仅是不需要判断是不是当前元素了，因为当前元素还没有添加进 hash 里。

```go
func TwoSum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}
```

## 结语

题目比较简单，毕竟暴力的方法也可以解决。唯一闪亮的点就是，时间复杂度从 O（n²）降为 O（n） 的时候，对 hash 的应用，有眼前一亮的感觉。



