func lengthOfLastWord(s string) int {
    n := len(s)
    count := 0
    var flag bool = false
    for i := n-1 ; i>=0; i-- {
            if s[i]!=' ' {
            count++
            flag = true
            } else {
               if flag == true {
                   break
               }
         }
    }
        return count
    }
    

