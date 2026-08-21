-- Write your query below
select
CASE
    when from_id > to_id THEN to_id
    ELSE 
    from_id
END as person1
,
CASE
    when from_id > to_id THEN from_id
    ELSE 
    to_id
END as person2,
count(*) as call_count,
sum(duration) as total_duration
from calls
group by person1, person2

