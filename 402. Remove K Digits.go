//https://leetcode.com/problems/remove-k-digits/description/


func removeKdigits(num string, k int) string {
	n := 0
	r := []byte(num)
    for _, c := range r {
        for ; n > 0 && r[n - 1] > c && k > 0; {
		    
            n--
            k--
        }
        
        if c != '0' || n > 0 {
            r[n] = c
            n++
        }
    }

    n -= k
    
    if n <= 0 {
        return "0"
    }
    
    return string(r[:n])
}
