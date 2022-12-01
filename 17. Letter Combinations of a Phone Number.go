// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

func letterCombinations(digits string) []string {
    
    ans := []string{}
    if len(digits) == 0 {
        return ans
    }
    helper(digits, 0, "", &ans)
    return ans
}


var M = map[byte]string {
    '1': "",
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
    '0': " ",
}

func helper(digits string, pos int, s string, strs *[]string) {
     if pos == len(digits) {
        *strs = append(*strs, s)
        return
    }
    
    for _, c := range M[digits[pos]] {
        helper(digits, pos + 1, s + string(c), strs)
    }
    return
}
