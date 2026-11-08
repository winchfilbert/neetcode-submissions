-- Write your query below
-- for the explanation of the case, if the team doesn't play then go with 0, 
-- that's why coalesce is used, if win then 3, if lose then 2


-- per row the visualization is gonna like
-- 10, the name of team, the sum will look rom every match, equals 7 for team id 10
-- etc
select t.team_id, t.team_name, 
    COALESCE(SUM(
        CASE
            WHEN t.team_id = m.host_team AND m.host_goals > m.guest_goals THEN 3
            WHEN t.team_id = m.guest_team AND m.guest_goals > m.host_goals THEN 3
            WHEN m.host_goals = m.guest_goals THEN 1
            ELSE 0
        END
    ), 0) AS num_points
from teams t
left join matches m 
ON t.team_id = m.host_team 
OR t.team_id = m.guest_team 
GROUP BY t.team_id, t.team_name
order by num_points DESC, team_id ASC

-- i just don't know why OR is used in here