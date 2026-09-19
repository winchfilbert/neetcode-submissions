-- Write your query below
-- what confuses me the most is that this is one table, but just represented with different alias, but this one not using group by tho
select a.sale_date, a.sold_num - o.sold_num as diff
from sales a join sales o on a.sale_date = o.sale_date
where a.fruit = 'apples' and o.fruit = 'oranges'
order by a.sale_date