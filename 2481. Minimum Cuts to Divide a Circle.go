// https://leetcode.com/problems/minimum-cuts-to-divide-a-circle/description/




func numberOfCuts(n int) int {
    
    if n == 1 {
        return 0
    } else if n%2==0 {
        return n/2
    } else {
        return n
    }
    
}
