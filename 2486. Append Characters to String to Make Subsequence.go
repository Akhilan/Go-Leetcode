// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/description/

func appendCharacters(s string, t string) int {
    i,j:=0,0
    for i<len(s)&&j<len(t){
        if t[j]==s[i]{
            i++
            j++
        }else{
            i++
        }
    }
    
    return len(t)-j


    
    
    
    
}
