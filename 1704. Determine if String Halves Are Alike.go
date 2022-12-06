// https://leetcode.com/problems/determine-if-string-halves-are-alike/description/


func halvesAreAlike(s string) bool {
    // s = strings.ToLower(s)
    mid := len(s)/2
    first := s[:mid]
    second := s[mid:]
    fir, sec := 0, 0
    
    for _, value := range first {
    switch value {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        fir++
    }
   }
    
    for _, value := range second {
    switch value {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        sec++
    }
   }
    
    if fir == sec {
        return true
    }
    
    
    return false
}





func halvesAreAlike(s string) bool {
    first := s[:len(s)/2]
    second := s[len(s)/2:]
  
    fir := vowelCount(first)
    sec := vowelCount(second) 
    
    if fir == sec {
        return true
    }
    return false
}


func vowelCount(s string) int {
    var count int
    for _, value := range s {
    switch value {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        count++
    }
   }
    return count
}
