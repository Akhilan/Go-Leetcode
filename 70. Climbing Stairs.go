// https://leetcode.com/problems/climbing-stairs/

//You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func climbStairs(n int) int {
    if n<=2 {
        return n
    }
    n1, n2 := 1, 2
    for i:=3; i<=n; i++{
       n1, n2 = n2, n1+n2
    }
    return n2
}
