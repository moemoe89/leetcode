SELECT s.name
FROM SalesPerson s
LEFT JOIN Orders o ON o.sales_id = s.sales_id
LEFT JOIN Company c ON c.com_id = o.com_id
GROUP BY s.sales_id
HAVING SUM(IF(c.name = 'RED', 1, 0)) = 0
