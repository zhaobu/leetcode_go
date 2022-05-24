--
-- @lc app=leetcode.cn id=197 lang=mysql
--
-- [197] 上升的温度
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1
 */
select a.id
from Weather a
    join Weather b on DATEDIFF(a.recordDate, b.recordDate) = 1
    and a.temperature > b.temperature;

-- @lc code=end