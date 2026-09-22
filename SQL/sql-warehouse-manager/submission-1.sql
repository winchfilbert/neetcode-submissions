-- 1. Select the warehouse name and calculate the total volume (width * length * height)
SELECT w.name AS warehouse_name, SUM(p.width * p.length * p.height * w.units) AS volume
-- 2. Start with the warehouse table (aliased as 'w')
FROM warehouse w 
-- 3. Connect it to the products table (aliased as 'p') where the product IDs match
JOIN products p ON w.product_id = p.product_id
-- 4. Group the results so you get one total volume calculated per unique warehouse name
GROUP BY w.name;