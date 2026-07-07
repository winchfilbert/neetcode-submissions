-- Write your query below
select email from person 
-- cara ngitung appearance pake group by and having, nice lesson
group by email
having count(email) > 1