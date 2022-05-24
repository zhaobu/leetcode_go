--
-- @lc app=leetcode.cn id=181 lang=mysql
--
-- [181] 超过经理收入的员工
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 使用2个表where
 */
select a.name as "Employee"
from Employee a,
    Employee b
where a.salary > b.salary
    and a.managerId = b.id;

/* 
 解法2 使用join语句
 */
select a.name as "Employee"
from Employee as a
    join Employee as b on a.salary > b.salary
    and a.managerId = b.id;

-- @lc code=end