-- Write your query below
select distinct(customer_id) from customers 
where revenue > 0 and customers.year = 2020