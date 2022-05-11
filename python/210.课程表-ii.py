#
# @lc app=leetcode.cn id=210 lang=python3
#
# [210] 课程表 II
#

# @lc code=start

from typing import List


class Solution:

    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        # 有向边
        edges = [[] for _ in range(numCourses)]
        # 节点的入度
        outDeg = [0] * numCourses
        # print(edges, outDeg)

        # 初始化有向边
        for v in prerequisites:
            start, end = v[1], v[0]
            edges[start].append(end)
            outDeg[end] += 1

        queue = [i for i, v in enumerate(outDeg) if v == 0]

        ret = []
        while len(queue) > 0:
            head = queue[0]
            queue = queue[1:]
            ret.append(head)
            for end in edges[head]:
                outDeg[end] -= 1
                if outDeg[end] == 0:
                    queue.append(end)

        if len(ret) != numCourses:
            return []
        return ret


# @lc code=end
