// https://leetcode.com/problems/circular-sentence/description/



func isCircularSentence(sentence string) bool {
    size := len(sentence) - 1
    
    for i := range sentence {
        
        if sentence[i] == ' ' {
            if sentence[i-1] != sentence[i+1] {
                return false
            }
        }
    }
    
    if sentence[0] == sentence[size] {
        return true
    }
    return false
}
