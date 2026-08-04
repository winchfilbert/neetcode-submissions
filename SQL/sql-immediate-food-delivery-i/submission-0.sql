-- Write your query below
SELECT 
    ROUND(
        SUM(CASE WHEN order_date = customer_pref_delivery_date THEN 1 ELSE 0 END) * 100.0 / COUNT(*), 
        2
    ) AS immediate_percentage 
FROM delivery;



-- this is the simpler version
-- explanations:
-- If you want to stick closer to your original structure without writing out long CASE WHEN statements, you can use boolean expressions. In many SQL dialects (like MySQL), a condition like order_date = customer_pref_delivery_date evaluates natively to 1 (true) or 0 (false), allowing you to just SUM them up directly.  
-- SELECT 
--     ROUND(
--         SUM(order_date = customer_pref_delivery_date) * 100.0 / COUNT(*), 
--         2
--     ) AS immediate_percentage 
-- FROM delivery;