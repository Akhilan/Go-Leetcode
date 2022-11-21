// https://leetcode.com/problems/palindrome-number/

// Given an integer x, return true if x is a palindrome, and false otherwise.


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var a []int
	for i := x; i != 0; i /= 10 {
		a = append(a, i%10)
	}
	j := int8(len(a) - 1)
	for _, v := range a {
		if v != a[j] {
			return false
		}
		j--
	}
	return true
}
