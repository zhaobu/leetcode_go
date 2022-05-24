--
-- @lc app=leetcode.cn id=184 lang=mysql
--
-- [184] 部门工资最高的员工
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1
 */
select b.name as Department,
    a.name as Employee,
    a.salary as Salary
from Employee a
    join Department b on (a.salary, a.departmentId) in (
        select max(salary) as salary,
            departmentId
        from Employee
        group by departmentId
    )
    and a.departmentId = b.id;

/* 
 解法2
 */
select b.name as Department,
    a.name as Employee,
    a.salary as Salary
from Employee a
    join Department b on a.departmentId = b.id
where (a.salary, a.departmentId) in (
        select max(salary) as salary,
            departmentId
        from Employee
        group by departmentId
    );

-- @lc code=end