--
-- @lc app=leetcode.cn id=177 lang=mysql
--
-- [177] 第N高的薪水
--
-- @lc code=start
create function getNthHighestSalary(N INT) returns INT begin
set N = N -1;

return (
  # Write your MySQL query statement below.
  select (
      select distinct salary
      from Employee
      order by salary desc
      limit 1 offset N
    )
);
end;

-- @lc code=end