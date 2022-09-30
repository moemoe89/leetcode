SELECT l1.num AS ConsecutiveNums
FROM Logs l1
LEFT JOIN Logs l2
ON l1.id + 1 = l2.id
LEFT JOIN Logs l3
ON l1.id - 1 = l3.id
WHERE l1.num = l2.num AND l1.num = l3.num
GROUP BY ConsecutiveNums
