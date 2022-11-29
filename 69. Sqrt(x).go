
//https://leetcode.com/problems/sqrtx/


//Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

func mySqrt(x int) int {
    for i := 1; i <= x; i++ {
        if i*i == x { 
            return i 
        }
        if i*i > x { 
            return i-1 
        }
    }
    return 0
}
