func isSubsequence(s string, t string) bool {
    found := 0
    for i,j:=0,0;i<len(t)&&j<len(s);i++ {
        if s[j] == t[i] {
            j++
            found++
        }
    }
    return found == len(s)
}
