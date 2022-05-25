--
-- @lc app=leetcode.cn id=584 lang=mysql
--
-- [584] 寻找用户推荐人
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 
 注意mysql比较是否为空要用is null不能用 =null 或者 !=null
 */
select name
from customer
where referee_id != 2
    or referee_id is null;

-- @lc code=end