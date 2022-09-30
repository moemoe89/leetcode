SELECT MAX(salary) AS SecondHighestSalary
FROM Employee
WHERE salary < (SELECT MAX(salary) FROM Employee)


SELECT MAX(salary) AS SecondHighestSalary
FROM Employee
WHERE salary < (SELECT DISTINCT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 0)
