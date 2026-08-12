-- Write your query below
select actor_id, director_id from actor_director
group by actor_id, director_id
-- remember that aggregate can be evaluated using only, not where
having count(timestamp) >= 3
