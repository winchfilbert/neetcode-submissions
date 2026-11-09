-- Write your query below

select min(ABS(p1.x - p2.x)) as shortest from point p1 join point p2 on p1.x < p2.x;

-- by doing the condition at the 'ON' operation
-- if the condition is true, therefore p1.x value is lesser
-- then the p2.x, so p1.x will always be the index and p2.x
-- will be next to that index
-- this conditions is fulfilled by the clue "sorted in ASC" 