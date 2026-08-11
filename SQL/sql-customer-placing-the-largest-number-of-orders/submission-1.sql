-- Write your query below
select customer_number from orders
group by customer_number
order by count(*) DESC
Limit 1