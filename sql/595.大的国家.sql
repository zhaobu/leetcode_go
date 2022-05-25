--
-- @lc app=leetcode.cn id=595 lang=mysql
--
-- [595] 大的国家
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 or
 */
select name,
    population,
    area
from World
where area >= 3000000
    or population >= 25000000;


/* 
 解法2 union
 */
select name,
    population,
    area
from World
where area >= 3000000
union
select name,
    population,
    area
from World
where population >= 25000000;

-- @lc code=end