// https://leetcode.com/problems/valid-parentheses/


// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

//    Open brackets must be closed by the same type of brackets.
//    Open brackets must be closed in the correct order.
//    Every close bracket has a corresponding open bracket of the same type.



func isValid(s string) bool {
    
    stack := []rune{}
    
    for _,val := range s{
        
        switch val{
            case '(','{','[':
                 stack = append(stack, val)
                    break
            
            
            case ')':
                 if len(stack) == 0 || stack[len(stack)-1] != '('{
                return false
                 }else{
                     stack = stack[:len(stack)-1]
                 }
            
            case ']':
                 if len(stack) == 0 || stack[len(stack)-1] != '['{
                return false
                 }else{
                     stack = stack[:len(stack)-1]
                 }
            

            
            case '}':
                 if len(stack) == 0 || stack[len(stack)-1] != '{'{
                return false
                 }else{
                     stack = stack[:len(stack)-1]
                 }
            
        }
    }
    
    return len(stack) == 0
}
