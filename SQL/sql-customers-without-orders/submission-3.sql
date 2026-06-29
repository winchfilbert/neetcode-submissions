-- Write your query below
-- Left Join: Grabs the entire left circle (all customers), 
-- regardless of whether they overlap with the right circle.
select customers.name
from customers 
left join orders on customers.id = orders.customer_id
where orders.customer_id iS NULL