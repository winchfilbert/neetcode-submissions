-- Write your query below
select employee_id, 
case when name not LIKE 'M%' and employee_id % 2 = 1 then salary
ELSE 0
end as bonus from employees 
order by employee_id ASC