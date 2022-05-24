--
-- @lc app=leetcode.cn id=185 lang=mysql
--
-- [185] 部门工资前三高的所有员工
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 使用子查询
 */
select a.name as Department,
    e1.name as Employee,
    e1.salary as Salary
from Employee as e1
    join Department as a on a.id = e1.departmentId
where 3 > (
        /* 
         1. 如果一个员工薪水在部门中排在前三名,则去重相同薪水后比他的薪水高的员工个数肯定<3
         2. 比如如果薪水在部门排第1.则同部门比他薪水高的员工为0
         如果薪水排名第3,则同部门比他薪水高的只能是第1和第2名,去重后最多只能有2人比他高
         */
        select count(distinct e2.salary)
        from Employee e2
        where e2.departmentId = e1.departmentId
            and e2.salary > e1.salary
    );

/* 
 解法2 使用窗口函数
 按题目意思,相同薪水名次相同,并且不占用排名名额,也就是,1,1,2,3,4,4,5这种
 应该使用 DENSE_RANK() OVER 函数
 */
select a.name as Department,
    b.name as Employee,
    b.salary as Salary
from Department as a
    join (
        select dense_rank() over(
                partition by departmentId
                order by salary desc
            ) as rankNum,
            name,
            departmentId,
            salary
        from Employee
    ) as b on a.id = b.departmentId
where b.rankNum <= 3;

-- @lc code=end