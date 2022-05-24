--
-- @lc app=leetcode.cn id=183 lang=mysql
--
-- [183] 从不订购的客户
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 not in
 */
select Name as Customers
from Customers
where Id not in (
        select CustomerId
        from Orders
    );

/* 
 解法2 左连接
 */
select a.Name as 'Customers'
from Customers a
    left join Orders b on a.Id = b.CustomerId
where b.CustomerId is null;

-- @lc code=end