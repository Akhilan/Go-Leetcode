func construct2DArray(original []int, m int, n int) [][]int {
    
     if len(original) != m * n { 
         return [][]int{} 
     }
    
    ans := make([][]int, m)
    index := 0
    
    for i := 0; i < m; i++ {
        ans[i] = make([]int, n)
        for j := 0; j < n; j++ {
            ans[i][j] = original[index]
            index++
        }
        }
    return ans
        
}
