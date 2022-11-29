// https://leetcode.com/problems/power-of-three/

//Given an integer n, return true if it is a power of three. Otherwise, return false.

// An integer n is a power of three, if there exists an integer x such that n == 3x.


func isPowerOfThree(n int) bool {
    
    if n == 1{
        return true
    }
    if n % 3 != 0 || n <= 0{
        return false
    }
    num := 1
    for num <= n {
        num *= 3
        if(num == n){
            return true
        }
    }
    return false
}
