SELECT * FROM employees

SELECT COUNT(1) FROM employees WHERE job_title = 'Manager' 

SELECT name, salary FROM employees WHERE department IN ('Sales','Marketing')

SELECT AVG(salary) AS average_salary
FROM employees
WHERE joined_date >= CURDATE() - INTERVAL 5 YEAR;

SELECT e.employee_id , e.name, SUM(s.sales) AS total_sales
FROM employees e
JOIN sales_data s ON e.employee_id = s.employee_id
GROUP BY e.employee_id , e.name
ORDER BY total_sales DESC
LIMIT 5;

WITH DepartmentAverage AS (
    SELECT department, AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department
),
OverallAverage AS (
    SELECT AVG(avg_salary) AS overall_avg_salary
    FROM DepartmentAverage
),
DepartmentsWithHigherAvg AS (
    SELECT department, avg_salary
    FROM DepartmentAverage
    JOIN OverallAverage
    ON avg_salary > overall_avg_salary
)
SELECT e.name, e.salary, d.avg_salary
FROM employees e
JOIN DepartmentsWithHigherAvg d
ON e.department = d.department
JOIN DepartmentAverage da
ON e.department = da.department

WITH TotalSales AS (
    SELECT e.employee_id, e.name, COALESCE(SUM(s.sales), 0) AS total_sales
    FROM employees e
    LEFT JOIN sales_data s ON e.employee_id = s.employee_id
    GROUP BY e.employee_id, e.name
),
RankedSales AS (
    SELECT employee_id, name, total_sales,
           ROW_NUMBER() OVER (ORDER BY total_sales DESC) AS rank
    FROM TotalSales
)
SELECT name, total_sales, rank
FROM RankedSales;

DELIMITER //
CREATE PROCEDURE GetEmployeesByDepartment(
    IN departmentName VARCHAR(255)
)
BEGIN
    SELECT e.name, SUM(e.salary) AS total_salary
    FROM employees e
    JOIN departments d ON e.department_id = d.id
    WHERE d.department_name = departmentName
    GROUP BY e.name;
END //
DELIMITER ;