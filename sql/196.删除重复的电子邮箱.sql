--
-- @lc app=leetcode.cn id=196 lang=mysql
--
-- [196] 删除重复的电子邮箱
--
-- @lc code=start
# Please write a DELETE statement and DO NOT write a SELECT statement.
# Write your MySQL query statement below
/* 
 解法1
 注意如果使用了表别名,则delet后面要写上别名
 */
delete p1
from Person p1,
    Person p2
where p1.email = p2.email
    and p1.id > p2.id;

-- @lc code=end