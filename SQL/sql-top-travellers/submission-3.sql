-- Write your query below

-- left join biar list of namesnya keliatan
select u.name, Coalesce(sum(r.distance), 0) as travelled_distance from users u left join rides r on u.id = r.user_id
group by u.id, u.name
order by travelled_distance DESC, u.name ASC -- you can order by per column