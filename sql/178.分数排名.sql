--
-- @lc app=leetcode.cn id=178 lang=mysql
--
-- [178] 分数排名
--
-- @lc code=start
# Write your MySQL query statement below

/* 
解法1 使用窗口函数
 */
select score,
    dense_rank() over (
        order by score desc
    ) as "rank" #这个rank之所以要加引号，因为rank本身是个函数，直接写rank会报错
from Scores;

-- @lc code=end