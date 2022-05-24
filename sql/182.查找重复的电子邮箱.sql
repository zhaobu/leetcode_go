--
-- @lc app=leetcode.cn id=182 lang=mysql
--
-- [182] 查找重复的电子邮箱
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 临时表
 */
select Email
from (
        select count(*) as num,
            Email
        from Person
        group by Email
    ) statictable
where num > 1;

/* 
 解法2 使用having
 */
select Email
from Person
group by Email
having count(*) > 1;

-- @lc code=end