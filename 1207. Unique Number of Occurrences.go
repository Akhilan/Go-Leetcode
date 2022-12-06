// https://leetcode.com/problems/unique-number-of-occurrences/description/


func uniqueOccurrences(arr []int) bool {
    
    freq := make(map[int]int)
    
    for _ , num :=  range arr {
        freq[num]++
    }
    
    var res []int
    
    for _, y := range freq {
        res = append(res, y)
    }
    
    
    sort.Ints(res)
    
    for k := 1; k < len(res); k++ {
        if res[k-1] == res[k] {
            return false
        }
    }
    
    return true
}
    
    
    
    
    
    


func uniqueOccurrences(arr []int) bool {
    
    freq := make(map[int]int)
    temp := make(map[int]int)
    
    for _ , num :=  range arr {
        freq[num]++
    }
    
    for _, v := range freq {
        temp[v]++
        if temp[v] > 1 { return false }
    }
    return true
    
    
}




