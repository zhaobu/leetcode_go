# [1. 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

## 问题

### 描述

> 编写一个函数来查找字符串数组中的最长公共前缀。
>
> 如果不存在公共前缀，返回空字符串 `""`。

### 示例 1

> ```
> 输入: ["flower","flow","flight"]
> 输出: "fl"
> ```

### 示例 2

> ```
> 输入: ["dog","racecar","car"]
> 输出: ""
> 解释: 输入不存在公共前缀。
> ```

### 说明

- 所有输入只包含小写字母 `a-z` 。



## 题解

### 思路 1 哈希表

> 首先拿到这道题，我们基本马上可以想到，此题可以看成是一道传统的**映射题（map 映射）**，为什么可以这样看呢，因为**我们需找出两个数组的交集元素，同时应与两个数组中出现的次数一致。这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>**。剩下的就是顺利成章的解题。

```go
func Intersect1(nums1 []int, nums2 []int) []int {
	nums0 := map[int]int{}
	for _, v := range nums1 {
		nums0[v] += 1
	}

	k := 0
	for _, v := range nums2 {
		if nums0[v] > 0 {
			nums0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}
```

## 复杂度分析

- 时间复杂度：O(m+n)O(m+n)，其中 mm 和 nn 分别是两个数组的长度。需要遍历两个数组并对哈希表进行操作，哈希表操作的时间复杂度是 O(1)O(1)，因此总时间复杂度与两个数组的长度和呈线性关系。
- 空间复杂度：O(\min(m,n))O(min(m,n))，其中 mm 和 nn 分别是两个数组的长度。对较短的数组进行哈希表的操作，哈希表的大小不会超过较短的数组的长度。为返回值创建一个数组 intersection，其长度为较短的数组的长度。

### 思路 2 排序

> 首先对两个数组进行排序，然后使用两个指针遍历两个数组。
>
> 初始时，两个指针分别指向两个数组的头部。每次比较两个指针指向的两个数组中的数字，如果两个数字不相等，则将指向较小数字的指针右移一位，如果两个数字相等，将该数字添加到答案，并将两个指针都右移一位。当至少有一个指针超出数组范围时，遍历结束。

```go
func Intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	res := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
```

## 复杂度分析

[leetcode 官方](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/solution/liang-ge-shu-zu-de-jiao-ji-ii-by-leetcode-solution/)

## 结语
