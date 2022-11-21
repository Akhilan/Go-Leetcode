// https://leetcode.com/problems/reverse-string/

func reverseString(s []byte)  {
    
    size := len(s) - 1
    
    //swapping elements
    for i := 0 ; i< size ; i++ {
        s[i], s[size] = s[size], s[i]
        size--
    }
}
