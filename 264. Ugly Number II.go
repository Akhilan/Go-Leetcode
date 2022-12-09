// https://leetcode.com/problems/ugly-number-ii/description/

// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

// Given an integer n, return the nth ugly number.

func nthUglyNumber(n int) int {
	if n == 1 {
		return n
	}
	fives, threes, twos := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(dp[twos]*2, min(dp[threes]*3, dp[fives]*5))
		if dp[i] == dp[twos]*2 {
			twos++
		}
		if dp[i] == dp[threes]*3 {
			threes++
		}
		if dp[i] == dp[fives]*5 {
			fives++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
