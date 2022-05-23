"""
# Definition for a Node.
"""


class Node:

    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:

    def treeToDoublyList(self, root: 'Node') -> 'Node':
        if not root:
            return root

        self.pre = None

        def dfs(node: Node):
            if not node:
                return

            dfs(node.left)
            if not self.pre:
                self.head = node
            else:   
                self.pre.right = node
                node.left = self.pre
            self.pre = node
            dfs(node.right)
            return

        dfs(root)
        # 最后再处理头节点的前驱和尾节点的后继
        self.pre.right = self.head
        self.head.left = self.pre
        return self.head