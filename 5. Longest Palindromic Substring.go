// https://leetcode.com/problems/longest-palindromic-substring/

// Given a string s, return the longest palindromic substring in s.


func longestPalindrome(s string) string {
    
    var max string
    for i := 0; i < len(s); i++ {
        max = maxPalindrome(s, i, i, max)
        max = maxPalindrome(s, i, i+1, max)
    }
    return max
}


func maxPalindrome(s string, i, j int, max string) string {
    var sub string
    for i >= 0 && j < len(s) && s[i] == s[j] {
        sub = s[i:j+1]
        i--
        j++
    }
    if len(max) < len(sub) {
        return sub
    }
    return max
    }


    

    
    
    

