// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
    
    if len(s) != len(t) {
        return false
    }

    mapper := make([]int, 26)

    for i:= range s {
        mapper[int(s[i]) - int('a')] += 1
        mapper[int(t[i]) - int('a')] -= 1
    }

   
    for _, v := range mapper {
        if v != 0 {
            return false
        }
    }
    return true
}
