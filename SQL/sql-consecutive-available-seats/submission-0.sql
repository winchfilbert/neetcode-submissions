-- Write your query below
SELECT DISTINCT c1.seat_id
FROM cinema c1
-- looking at left with c1.seat_id = c2.seat_id + 1 and right with c1.seat_id = c2.seat_id -1)
-- therefore it is ensure that at least 2 seats are available
JOIN cinema c2
    ON (c1.seat_id = c2.seat_id + 1 OR c1.seat_id = c2.seat_id - 1)
WHERE c1.free = 1 AND c2.free = 1
ORDER BY c1.seat_id;
