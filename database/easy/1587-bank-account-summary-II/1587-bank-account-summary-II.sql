SELECT
    name AS NAME,
    SUM(amount) AS BALANCE
FROM Users u
JOIN Transactions t ON u.account = t.account
GROUP BY u.account
HAVING BALANCE > 10000
