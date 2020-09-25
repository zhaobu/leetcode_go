# [1. 两个数组的交集 II][https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/]

## 问题

### 描述

> 给定两个数组，编写一个函数来计算它们的交集。

### 示例 1

> 输入：nums1 = [1,2,2,1], nums2 = [2,2]
> 输出：[2,2]

### 示例 2

> 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
> 输出：[4,9]

### 说明

- 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
- 我们可以不考虑输出结果的顺序。

### 进阶

- 如果给定的数组已经排好序呢？你将如何优化你的算法？
- 如果  nums1  的大小比  nums2  小很多，哪种方法更优？
- 如果  nums2  的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

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
