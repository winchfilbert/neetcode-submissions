-- Write your query below
select distinct(employee_id), 
CASE 
        WHEN name NOT LIKE 'M%' AND employee_id % 2 = 1 THEN salary
        ELSE 0 
END AS bonus from employees 
ORDER BY employee_id ASC