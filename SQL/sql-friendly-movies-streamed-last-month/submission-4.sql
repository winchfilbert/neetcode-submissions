-- Write your query below
select distinct c.title from content c join tv_program t on t.content_id = c.content_id where c.kids_content = 'Y' and c.content_type = 'Movies' 
intersect
select distinct c.title from content c join tv_program t on t.content_id = c.content_id 
where t.program_date >= '2020-06-01 00:00' and t.program_date <= '2020-06-30 23:59';