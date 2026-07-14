-- Write your query below
select p.first_name, p.last_name, coalesce(a.city, NULL) AS city, coalesce(a.state, NULL) as state from person p left join address a on p.person_id = a.person_id