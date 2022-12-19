// https://leetcode.com/problems/backspace-string-compare/description/

func backspaceCompare(S string, T string) bool {
	return helper(S) == helper(T)
}
func helper(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if s[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
