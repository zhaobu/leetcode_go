#
# @lc app=leetcode.cn id=40 lang=python3
#
# [40] 组合总和 II
#

# @lc code=start
from typing import List


class Solution:

    def dfs(self, start: int, sum: int, record: List):
        if sum == self.target:
            # print("record={}".format(record))
            # 最终结果要复制
            self.ret.append(record.copy())
            return
        for i in range(start, self.n):
            # 确保在一种结果当中,下标相同的位置,元素不重复,就能保证最终结果没有重复解
            if i > start and self.candidates[i] == self.candidates[i - 1]:
                continue
            # 剪枝已经不能满足条件的情况
            if sum + self.candidates[i] > self.target:
                continue
            record.append(self.candidates[i])
            self.dfs(i + 1, sum + self.candidates[i], record)
            record.pop()

    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        self.n = len(candidates)
        self.candidates = candidates
        self.target = target
        self.ret = []  #保存结果
        self.candidates.sort()  #先排序,使相同数字能挨在一起
        self.dfs(0, 0, [])
        return self.ret


# @lc code=end
