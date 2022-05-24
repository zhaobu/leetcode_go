--
-- @lc app=leetcode.cn id=180 lang=mysql
--
-- [180] 连续出现的数字
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 3表关联查询
 */
select distinct l1.Num as ConsecutiveNums
from logs as l1,
    logs as l2,
    logs as l3
where l1.Num = l2.Num
    and l2.Num = l3.Num
    and l1.Id + 1 = l2.Id
    and l2.Id + 1 = l3.Id;

-- @lc code=end