#
# @lc app=leetcode.cn id=912 lang=python3
#
# [912] 排序数组
#

# @lc code=start
import random
from typing import List


class Solution1:
    """ 解法1 堆排序 """

    def heapify(self, nums, i, n):
        """  自顶向下调整堆 """
        while i < n:
            maxIdx = i
            left, right = 2 * i + 1, 2 * i + 2
            if left < n and nums[left] > nums[maxIdx]:
                maxIdx = left
            if right < n and nums[right] > nums[maxIdx]:
                maxIdx = right
            if maxIdx == i:
                break
            nums[maxIdx], nums[i] = nums[i], nums[maxIdx]
            i = maxIdx

    def sortArray(self, nums: List[int]) -> List[int]:
        """
        1. 从最后一个非叶子节点开始堆化
              0
          1      2
        3   4  5   6
        """
        i = (len(nums) >> 1) - 1
        while i >= 0:
            self.heapify(nums, i, len(nums))
            i -= 1
        # 每次都把堆顶元素和最后一个元素交换,然后重新调整堆
        i = len(nums)
        while i > 0:
            i -= 1
            # print("1  nums={}".format(nums))
            nums[0], nums[i] = nums[i], nums[0]
            # print("2  nums={}".format(nums))
            self.heapify(nums, 0, i)
            # print("3  nums={}\n".format(nums))

        return nums


class Solution:
    """ 解法2 快速排序 """

    def sortArray(self, nums: List[int]) -> List[int]:
        self.sort(nums, 0, len(nums) - 1)
        return nums

    def sort(self, nums, i, j):
        if i >= j:
            return
        k = self.partition(nums, i, j)
        self.sort(nums, i, k - 1)
        self.sort(nums, k + 1, j)

    def partition(self, nums, i, j) -> int:
        k = random.randint(i, j)
        nums[j], nums[k] = nums[k], nums[j]
        last = i  # last表示第一个大于nums[j]的元素下标
        # print("i={},j={},k={} ".format(i, j, k))
        for cur in range(i, j):
            # print("i={},j={},cur={}".format(i, j, cur))
            if nums[cur] <= nums[j]:
                nums[cur], nums[last] = nums[last], nums[cur]
                last += 1
        nums[last], nums[j] = nums[j], nums[last]
        return last


# @lc code=end