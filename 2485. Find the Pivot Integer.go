// https://leetcode.com/problems/find-the-pivot-integer/description/


func pivotInteger(n int) int {
    upSum := 0
    for i:=n ; i > n/2 ; i-- {
        upSum +=i
        downSum := i*(i+1)/2
        if upSum == downSum {
            return i
        }
    }
    return -1
}
