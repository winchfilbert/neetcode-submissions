-- Write your query below
select s.seller_name from seller s left join orders o on s.seller_id = o.seller_id and Extract(YEAR from o.sale_date) = 2020 
GROUP BY s.seller_id, s.seller_name
having count(o.seller_id) = 0
order by s.seller_name ASC