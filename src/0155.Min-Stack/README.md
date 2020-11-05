#### [最小栈](https://leetcode-cn.com/problems/min-stack/)

## 问题

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

**示例 1:**

```
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

**提示：**

- `pop`、`top` 和 `getMin` 操作总是在 **非空栈** 上调用。

## 题解

1. [windliang](https://leetcode.wang/leetcode100%E6%96%A9%E5%9B%9E%E9%A1%BE.html)
2. [leetcode 官方](https://leetcode-cn.com/problemset/all/)
3. [小浩算法](https://www.geekxh.com/0.0.%E5%AD%A6%E4%B9%A0%E9%A1%BB%E7%9F%A5/01.html)

## 结语
