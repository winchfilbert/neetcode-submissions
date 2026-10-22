-- Write your query below

-- this still cause duplicates
-- select e.student_id, e.exam_id, e.score from exam_results e INNER JOIN (
--     select student_id, max(score) as score from exam_results group by student_id
-- ) highest on e.student_id = highest.student_id AND e.score = highest.score
-- order by e.student_id ASC, e.score DESC


-- this works, because it's isolate only to one exam, preventing if there's two id with same score
SELECT e.student_id, e.exam_id, e.score
FROM exam_results e
WHERE e.exam_id = (
    SELECT e2.exam_id
    FROM exam_results e2
    WHERE e2.student_id = e.student_id
    ORDER BY e2.score DESC, e2.exam_id ASC
    LIMIT 1
)
ORDER BY e.student_id ASC;