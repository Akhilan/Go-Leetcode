# https://leetcode.com/problems/combine-two-tables/

SELECT p.firstName, p.lastName, a.city, a.state FROM Person p
LEFT JOIN Address a USING (personId)
