// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/description/

func minAddToMakeValid(s string) int { 
    left, right := 0, 0
    
    for _, i := range s {
        if i == '(' {
            left++
        } else if i == ')' && left > 0 {
            left--
        } else {
            right++
        }
    }
 return left + right
}
