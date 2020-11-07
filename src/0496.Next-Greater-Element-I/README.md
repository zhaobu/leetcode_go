#### [下一个更大元素 I](https://leetcode-cn.com/problems/next-greater-element-i/)

## 问题

给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中 nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

### **示例 1：**

> ```text
> 输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
> 输出: [-1,3,-1]
> 解释:
>     对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
>     对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
>     对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
> ```

示例 2:

> ```
> 输入: nums1 = [2,4], nums2 = [1,2,3,4].
> 输出: [3,-1]
> 解释:
>     对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
>     对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
> ```

**提示：**

1. `nums1`和`nums2`中所有元素是唯一的。
2. `nums1`和`nums2` 的数组大小都不超过 1000。

## 题解

1. [windliang](https://leetcode.wang/leetcode100%E6%96%A9%E5%9B%9E%E9%A1%BE.html)
2. [leetcode 官方](https://leetcode-cn.com/problemset/all/)
3. [小浩算法](https://www.geekxh.com/0.0.%E5%AD%A6%E4%B9%A0%E9%A1%BB%E7%9F%A5/01.html)

## 结语
