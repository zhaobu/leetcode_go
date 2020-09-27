# [两个数组的交集 II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)

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

[leetcode官方](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/solution/liang-ge-shu-zu-de-jiao-ji-ii-by-leetcode-solution/)

## 结语
