--
-- @lc app=leetcode.cn id=176 lang=mysql
--
-- [176] 第二高的薪水
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 临时表
 */
select (
        select distinct salary
        from Employee
        order by salary desc
        limit 1 offset 1
    ) as SecondHighestSalary;

/*
 解法2 使用 IFNULL 和 LIMIT 子句
 */
select ifnull(
        (
            select distinct salary
            from Employee
            order by salary desc
            limit 1 offset 1
        ),
        null
    ) as SecondHighestSalary;

-- @lc code=end