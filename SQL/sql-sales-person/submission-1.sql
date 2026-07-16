-- Write your query below
-- this is definitely a trick, i naively tried to join three table simulateneously without considering that you can except
select s.name from sales_person s where s.sales_id NOT in (
    select o.sales_id from orders o join company c on o.com_id = c.com_id where c.name = 'CRIMSON'
)


-- the way we do "except" in sql, can be by matching a column not into a query, that's the easiest way i think