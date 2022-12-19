// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

func evalRPN(tokens []string) int {
    stk := make([]int, 0)
    
    for _, val := range tokens {
        if val == "+" || val == "-" || val == "*" || val == "/" {
            
            first := stk[len(stk) - 2]
            second := stk[len(stk) - 1]
            
            stk = stk[: len(stk) - 2]
            
            stk = append(stk, op(rune(val[0]), first, second))
            
        } else {
            v, _ := strconv.Atoi(val)
            stk = append(stk, v)
        }
    }
    
    return stk[0]
}


func op(sym rune, first, second int) int {
    if sym == '+' {
        return first + second
    } else if sym == '-' {
        return first - second
    } else if sym == '*' {
        return first * second
    } else {return int(first / second)}
}
