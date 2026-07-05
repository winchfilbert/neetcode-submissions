-- Write your query below
select customers.customer_id, customers.customer_name from customers join orders on customers.customer_id = orders.customer_id
where orders.product_name = 'A' 

INTERSECT

select customers.customer_id, customers.customer_name from customers join orders on customers.customer_id = orders.customer_id
where orders.product_name = 'B'

EXCEPT

select customers.customer_id, customers.customer_name from customers join orders on customers.customer_id = orders.customer_id
where orders.product_name = 'C'

Order by customer_name ASC;

-- this is very smart actually, checking if true by see if it's match the list
-- SELECT c.customer_id, c.customer_name
-- FROM customers c
-- WHERE c.customer_id IN (
--     SELECT customer_id FROM orders WHERE product_name = 'A'
-- )
-- AND c.customer_id IN (
--     SELECT customer_id FROM orders WHERE product_name = 'B'
-- )
-- AND c.customer_id NOT IN (
--     SELECT customer_id FROM orders WHERE product_name = 'C'
-- )
-- ORDER BY c.customer_name;