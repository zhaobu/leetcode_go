--
-- @lc app=leetcode.cn id=511 lang=mysql
--
-- [511] 游戏玩法分析 I
--
-- @lc code=start
# Write your MySQL query statement below
/* 
 解法1 group by后使用聚合函数min
 */
select player_id,
    min(event_date) as first_login
from Activity
group by player_id;

/* 
 解法2 子查询
 */
select a.player_id,
    a.event_date as first_login
from (
        select player_id,
            event_date
        from Activity
        order by event_date asc
    ) as a
group by a.player_id
order by a.event_date asc;

-- @lc code=end