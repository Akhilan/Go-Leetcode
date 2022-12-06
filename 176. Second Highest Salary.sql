# https://leetcode.com/problems/second-highest-salary/description/


Select MAX(Salary) as SecondHighestSalary
 from Employee
where Salary < (Select MAX(Salary) from Employee)
